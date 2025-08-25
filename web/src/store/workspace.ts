import { uniqBy } from "lodash-es";
import { makeAutoObservable } from "mobx";
import { workspaceServiceClient } from "@/grpcweb";
import { workspaceServiceClient as workspaceServiceClientV2  } from "@/grpc";
import {
  WorkspaceProfile,
  WorkspaceSetting_Key,
  WorkspaceSetting,
  WorkspaceProfileSchema,
  WorkspaceSettingSchema,
  WorkspaceSetting_GeneralSettingSchema,
  WorkspaceSetting_MemoRelatedSettingSchema,
} from "@/types/proto2/model/workspace_pb";
import { create } from "@bufbuild/protobuf";
import { isValidateLocale } from "@/utils/i18n";
import { workspaceSettingNamePrefix } from "./common";

class LocalState {
  locale: string = "en";
  appearance: string = "system";
  profile: WorkspaceProfile = create(WorkspaceProfileSchema, {});
  settings: WorkspaceSetting[] = [];

  get generalSetting() {
    const setting = this.settings.find((setting) => setting.name === `${workspaceSettingNamePrefix}${WorkspaceSetting_Key.GENERAL}`);
    if (setting?.value?.case === "generalSetting") {
      return setting.value.value;
    }
    return create(WorkspaceSetting_GeneralSettingSchema, {});
  }

  get memoRelatedSetting() {
    const setting = this.settings.find((setting) => setting.name === `${workspaceSettingNamePrefix}${WorkspaceSetting_Key.MEMO_RELATED}`);

    if (setting?.value?.case === "memoRelatedSetting") {
      return setting.value.value;
    }
    return create(WorkspaceSetting_MemoRelatedSettingSchema, {});
  }

  constructor() {
    makeAutoObservable(this);
  }

  setPartial(partial: Partial<LocalState>) {
    const finalState = {
      ...this,
      ...partial,
    };
    if (!isValidateLocale(finalState.locale)) {
      finalState.locale = "en";
    }
    if (!["system", "light", "dark"].includes(finalState.appearance)) {
      finalState.appearance = "system";
    }
    Object.assign(this, finalState);
  }
}

const workspaceStore = (() => {
  const state = new LocalState();

  const fetchWorkspaceSetting = async (settingKey: WorkspaceSetting_Key) => {
    const setting = await workspaceServiceClientV2.getWorkspaceSetting({ name: `${workspaceSettingNamePrefix}${settingKey}` });
    state.setPartial({
      settings: uniqBy([setting, ...state.settings], "name"),
    });
  };

  const upsertWorkspaceSetting = async (setting: WorkspaceSetting) => {
    await workspaceServiceClient.updateWorkspaceSetting({ setting });
    state.setPartial({
      settings: uniqBy([setting, ...state.settings], "name"),
    });
  };

  const getWorkspaceSettingByKey = (settingKey: WorkspaceSetting_Key) => {
    return (
      state.settings.find((setting) => setting.name === `${workspaceSettingNamePrefix}${settingKey}`) || create(WorkspaceSettingSchema, {})
    );
  };

  return {
    state,
    fetchWorkspaceSetting,
    upsertWorkspaceSetting,
    getWorkspaceSettingByKey,
  };
})();

export const initialWorkspaceStore = async () => {
  const workspaceProfile = await workspaceServiceClientV2.getWorkspaceProfile({});
  // Prepare workspace settings.
  for (const key of [WorkspaceSetting_Key.GENERAL, WorkspaceSetting_Key.MEMO_RELATED]) {
    await workspaceStore.fetchWorkspaceSetting(key);
  }

  const workspaceGeneralSetting = workspaceStore.state.generalSetting;
  workspaceStore.state.setPartial({
    locale: workspaceGeneralSetting.customProfile?.locale,
    appearance: workspaceGeneralSetting.customProfile?.appearance,
    profile: workspaceProfile,
  });
};

export default workspaceStore;
