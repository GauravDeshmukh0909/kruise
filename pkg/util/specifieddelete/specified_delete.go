/*
Copyright 2020 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package specifieddelete

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
)

func IsSpecifiedDelete(obj metav1.Object) bool {
	_, ok := obj.GetLabels()[appsv1alpha1.SpecifiedDeleteKey]
	return ok
}

func PatchPodSpecifiedDelete(c client.Client, pod *v1.Pod, value string) (bool, error) {
	if _, ok := pod.Labels[appsv1alpha1.SpecifiedDeleteKey]; ok {
		return false, nil
	}

	body := fmt.Sprintf(
		`{"metadata":{"labels":{"%s":"%s"}}}`,
		appsv1alpha1.SpecifiedDeleteKey,
		value,
	)
	return true, c.Patch(context.TODO(), pod, client.RawPatch(types.StrategicMergePatchType, []byte(body)))
}
