/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	dataprotectionv1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	extensionsv1alpha1 "github.com/apecloud/kubeblocks/apis/extensions/v1alpha1"
	storagev1alpha1 "github.com/apecloud/kubeblocks/apis/storage/v1alpha1"
	workloadsv1alpha1 "github.com/apecloud/kubeblocks/apis/workloads/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.kubeblocks.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("backuppolicytemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().BackupPolicyTemplates().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().Clusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ClusterDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ClusterVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("components"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().Components().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("componentclassdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ComponentClassDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("componentdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ComponentDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("componentresourceconstraints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ComponentResourceConstraints().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("configconstraints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ConfigConstraints().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("opsdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().OpsDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("opsrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().OpsRequests().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("servicedescriptors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ServiceDescriptors().Informer()}, nil

		// Group=dataprotection.kubeblocks.io, Version=v1alpha1
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("actionsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().ActionSets().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().Backups().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backuppolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().BackupPolicies().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backuprepos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().BackupRepos().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backupschedules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().BackupSchedules().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("restores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().Restores().Informer()}, nil

		// Group=extensions.kubeblocks.io, Version=v1alpha1
	case extensionsv1alpha1.SchemeGroupVersion.WithResource("addons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Addons().Informer()}, nil

		// Group=storage.kubeblocks.io, Version=v1alpha1
	case storagev1alpha1.SchemeGroupVersion.WithResource("storageproviders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha1().StorageProviders().Informer()}, nil

		// Group=workloads.kubeblocks.io, Version=v1alpha1
	case workloadsv1alpha1.SchemeGroupVersion.WithResource("replicatedstatemachines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Workloads().V1alpha1().ReplicatedStateMachines().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
