<template>
  <div v-if="state.ready" class="w-full space-y-6">
    <DeploymentMatrix
      v-if="state.deployment"
      class="w-full !px-0 overflow-x-auto"
      :project="project"
      :deployment="state.deployment"
      :database-list="databaseList"
      :environment-list="environmentList"
      show-search-box
    />

    <div class="space-y-4">
      <div class="w-full flex flex-row justify-between items-center gap-4">
        <h3 class="text-lg font-medium leading-7 text-main">
          {{ $t("common.deployment-config") }}
        </h3>
        <div>
          <NButton v-if="allowEdit" @click="addStage">
            <template #icon>
              <PlusIcon class="w-4 h-auto" />
            </template>
            {{ $t("deployment-config.add-stage") }}
          </NButton>
        </div>
      </div>

      <BBAttention
        v-if="state.deployment === undefined"
        type="warning"
        :title="$t('common.deployment-config')"
        :description="$t('deployment-config.this-is-example-deployment-config')"
      >
      </BBAttention>
      <div v-else>
        <DeploymentConfigTool
          v-if="state.deployment.schedule"
          :schedule="state.deployment.schedule"
          :allow-edit="allowEdit"
          :database-list="databaseList"
        />
        <div class="pt-4 border-t flex justify-between items-center">
          <div class="flex items-center space-x-2">
            <NButton
              :disabled="!allowResetToDefaultDeploymentConfig"
              text
              @click="resetToDefaultDeploymentConfig"
            >
              <template #icon>
                <Undo2Icon class="w-4 h-auto" />
              </template>
              {{ $t("common.reset") }}
            </NButton>
          </div>
          <div v-if="allowEdit" class="flex items-center space-x-2">
            <NPopover v-if="allowEdit" :disabled="!state.error" trigger="hover">
              <template #trigger>
                <NButton
                  type="primary"
                  :disabled="!allowUpdateDeploymentConfig"
                  @click="updateDeploymentConfig"
                >
                  {{ $t("common.update") }}
                </NButton>
              </template>

              <span v-if="state.error" class="text-error">
                {{ $t(state.error) }}
              </span>
            </NPopover>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else class="flex justify-center items-center py-10">
    <BBSpin />
  </div>
</template>

<script lang="ts" setup>
import { cloneDeep, isEqual } from "lodash-es";
import { PlusIcon } from "lucide-vue-next";
import { Undo2Icon } from "lucide-vue-next";
import { NButton, NPopover, useDialog } from "naive-ui";
import type { PropType } from "vue";
import { computed, nextTick, reactive, watch, watchEffect } from "vue";
import { useI18n } from "vue-i18n";
import {
  getDefaultDeploymentConfig,
  pushNotification,
  useDeploymentConfigV1Store,
  useEnvironmentV1List,
} from "@/store";
import {
  DeploymentConfig,
  type LabelSelectorRequirement,
  type Project,
} from "@/types/proto/v1/project_service";
import { OperatorType } from "@/types/proto/v1/project_service";
import type { ComposedDatabase } from "../types";
import {
  extractEnvironmentResourceName,
  validateDeploymentConfigV1,
} from "../utils";
import DeploymentConfigTool, { DeploymentMatrix } from "./DeploymentConfigTool";
import { BBAttention, BBSpin } from "@/bbkit";

type LocalState = {
  ready: boolean;
  deployment: DeploymentConfig | undefined;
  originalDeployment: DeploymentConfig | undefined;
  error: string | undefined;
};

const props = defineProps({
  project: {
    required: true,
    type: Object as PropType<Project>,
  },
  databaseList: {
    type: Array as PropType<ComposedDatabase[]>,
    default: () => [],
  },
  allowEdit: {
    default: true,
    type: Boolean,
  },
});

const { t } = useI18n();
const dialog = useDialog();
const deploymentConfigV1Store = useDeploymentConfigV1Store();

const state = reactive<LocalState>({
  ready: false,
  deployment: undefined,
  originalDeployment: undefined,
  error: undefined,
});

const isDeploymentConfigDirty = computed((): boolean => {
  return !isEqual(state.deployment, state.originalDeployment);
});

const defaultDeploymentConfigSchedule = computed(
  () => getDefaultDeploymentConfig().schedule
);

const allowResetToDefaultDeploymentConfig = computed((): boolean => {
  return !isEqual(
    defaultDeploymentConfigSchedule.value,
    state.deployment?.schedule
  );
});

const allowUpdateDeploymentConfig = computed((): boolean => {
  if (state.error) return false;
  if (!isDeploymentConfigDirty.value) return false;
  return true;
});

const environmentList = useEnvironmentV1List();

const resetStates = async () => {
  await nextTick(); // Waiting for all watchers done
  state.error = undefined;
};

watchEffect(async () => {
  const deploymentConfig =
    await deploymentConfigV1Store.fetchDeploymentConfigByProjectName(
      props.project.name
    );
  // We clone the saved deployment-config
  // <DeploymentConfigTool /> will mutate `state.deployment` directly
  // when update button clicked, we save the draft to backend.
  state.deployment = cloneDeep(deploymentConfig);
  // clone the object to the backup
  state.originalDeployment = cloneDeep(state.deployment);
  // clean up error and dirty status
  resetStates();
  state.ready = true;
});

const addStage = () => {
  if (!state.deployment) return;
  const rule: LabelSelectorRequirement = {
    key: "environment",
    operator: OperatorType.OPERATOR_TYPE_IN,
    values: [],
  };
  if (environmentList.value.length > 0) {
    const name = extractEnvironmentResourceName(environmentList.value[0].name);
    rule.values.push(name);
  }

  state.deployment.schedule?.deployments.push({
    title: "New Stage",
    id: "",
    spec: {
      labelSelector: {
        matchExpressions: [rule],
      },
    },
  });
};

const validate = () => {
  if (!state.deployment) return;
  state.error = validateDeploymentConfigV1(state.deployment);
};

const resetToDefaultDeploymentConfig = () => {
  dialog.create({
    positiveText: t("common.confirm"),
    negativeText: t("common.cancel"),
    title: t("deployment-config.confirm-to-reset"),
    autoFocus: false,
    closable: false,
    maskClosable: false,
    closeOnEsc: false,
    onNegativeClick: () => {
      // nothing to do
    },
    onPositiveClick: () => {
      state.deployment = DeploymentConfig.fromPartial({
        ...cloneDeep(state.originalDeployment),
        schedule: getDefaultDeploymentConfig().schedule,
      });
      resetStates();
    },
  });
};

const updateDeploymentConfig = async () => {
  if (!state.deployment) return;
  if (!allowUpdateDeploymentConfig.value) return;

  const updated =
    await deploymentConfigV1Store.updatedDeploymentConfigByProjectName(
      props.project.name,
      state.deployment
    );

  pushNotification({
    module: "bytebase",
    style: "SUCCESS",
    title: t("deployment-config.update-success"),
  });

  // clone the updated version to the backup
  state.deployment = cloneDeep(updated);
  state.originalDeployment = cloneDeep(updated);
  // clean up error status
  resetStates();
};

watch(
  () => state.deployment,
  () => {
    validate();
  },
  { deep: true }
);
</script>
