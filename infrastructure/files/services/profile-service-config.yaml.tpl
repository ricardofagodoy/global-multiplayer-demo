# Copyright 2023 Google LLC All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: v1
kind: ServiceAccount
metadata:
  name: profile
  annotations:
    iam.gke.io/gcp-service-account: ${service_email}

---

apiVersion: v1
kind: ConfigMap
metadata:
  name: spanner-config
data:
  SPANNER_PROJECT_ID: ${project_id}
  SPANNER_INSTANCE_ID: ${instance_id}
  SPANNER_DATABASE_ID: ${database_id}