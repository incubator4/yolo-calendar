import { defineStore } from "pinia";
import request from "@/tools/axios";
import type { DescribeInstancesTrafficPackagesResponse } from "tencentcloud-sdk-nodejs/tencentcloud/services/lighthouse/v20200324/lighthouse_models";

export const useStatusStore = defineStore("status", () => {
  const traffic = ref<DescribeInstancesTrafficPackagesResponse>();
  const getTraffic = () =>
    request
      .get<{ Response: DescribeInstancesTrafficPackagesResponse }>(
        "/status/traffic"
      )
      .then((resp) => {
        traffic.value = resp.data.Response;
      });

  return {
    getTraffic,
    traffic,
  };
});
