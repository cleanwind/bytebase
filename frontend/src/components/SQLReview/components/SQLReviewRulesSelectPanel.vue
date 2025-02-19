<template>
  <Drawer :show="show" @close="$emit('close')">
    <DrawerContent
      :title="$t('sql-review.select-review')"
      class="w-[70rem] max-w-[100vw] relative"
    >
      <template #default>
        <SQLReviewTabsByEngine :rule-map-by-engine="ruleTemplateMapV2">
          <template
            #default="{
              ruleList: ruleListFilteredByEngine,
              engine,
            }: {
              ruleList: RuleTemplateV2[];
              engine: Engine;
            }"
          >
            <SQLRuleTableWithFilter
              :engine="engine"
              :rule-list="ruleListFilteredByEngine"
              :editable="false"
              :hide-level="true"
              :select-rule="true"
              :size="'small'"
              :selected-rule-keys="selectedRuleKeys"
              @update:selected-rule-keys="onSelectedRuleKeysUpdate"
            />
          </template>
        </SQLReviewTabsByEngine>
      </template>
      <template #footer>
        <div class="flex items-center justify-end gap-x-3">
          <NButton @click="$emit('close')">
            {{ $t("common.close") }}
          </NButton>
        </div>
      </template>
    </DrawerContent>
  </Drawer>
</template>

<script setup lang="ts">
import { NButton } from "naive-ui";
import { computed } from "vue";
import { Drawer, DrawerContent } from "@/components/v2";
import { type Engine, engineFromJSON } from "@/types/proto/v1/common";
import { ruleTemplateMapV2 } from "@/types/sqlReview";
import type { RuleTemplateV2 } from "@/types/sqlReview";
import SQLReviewTabsByEngine from "./SQLReviewTabsByEngine.vue";
import SQLRuleTableWithFilter from "./SQLRuleTableWithFilter.vue";
import { getRuleKey } from "./utils";

const props = defineProps<{
  show: boolean;
  selectedRuleMap: Map<Engine, Map<string, RuleTemplateV2>>;
}>();

const emit = defineEmits<{
  (event: "close"): void;
  (event: "rule-select", rule: RuleTemplateV2): void;
  (event: "rule-remove", rule: RuleTemplateV2): void;
}>();

const selectedRuleKeys = computed(() => {
  const keys = [];
  for (const map of props.selectedRuleMap.values()) {
    for (const rule of map.values()) {
      keys.push(getRuleKey(rule));
    }
  }
  return keys;
});

const onSelectedRuleKeysUpdate = (keys: string[]) => {
  const oldSelectedKeys = new Set([...selectedRuleKeys.value]);
  const newSelectedKeys = new Set(keys);

  for (const key of newSelectedKeys) {
    if (oldSelectedKeys.has(key)) {
      oldSelectedKeys.delete(key);
      continue;
    }
    const [engineStr, ruleKey] = key.split(":");
    const engine = engineFromJSON(engineStr);
    const rule = ruleTemplateMapV2.get(engine)?.get(ruleKey);
    if (rule) {
      emit("rule-select", rule);
    }
  }

  // keys remained in the oldSelectedKeys is not selected.
  for (const oldKey of oldSelectedKeys) {
    const [engineStr, ruleKey] = oldKey.split(":");
    const engine = engineFromJSON(engineStr);
    const rule = props.selectedRuleMap.get(engine)?.get(ruleKey);
    if (rule) {
      emit("rule-remove", rule);
    }
  }
};
</script>
