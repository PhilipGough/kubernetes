/*
Copyright The Kubernetes Authors.

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

package internalversion

import (
	context "context"
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	apps "k8s.io/kubernetes/pkg/apis/apps"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
	batch "k8s.io/kubernetes/pkg/apis/batch"
	certificates "k8s.io/kubernetes/pkg/apis/certificates"
	coordination "k8s.io/kubernetes/pkg/apis/coordination"
	core "k8s.io/kubernetes/pkg/apis/core"
	extensions "k8s.io/kubernetes/pkg/apis/extensions"
	networking "k8s.io/kubernetes/pkg/apis/networking"
	policy "k8s.io/kubernetes/pkg/apis/policy"
	rbac "k8s.io/kubernetes/pkg/apis/rbac"
	scheduling "k8s.io/kubernetes/pkg/apis/scheduling"
	settings "k8s.io/kubernetes/pkg/apis/settings"
	storage "k8s.io/kubernetes/pkg/apis/storage"
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
func (f *sharedInformerFactory) ForResource(ctx context.Context, resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=admissionregistration.k8s.io, Version=internalVersion
	case admissionregistration.SchemeGroupVersion.WithResource("initializerconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().InternalVersion().InitializerConfigurations().Informer(ctx)}, nil
	case admissionregistration.SchemeGroupVersion.WithResource("mutatingwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().InternalVersion().MutatingWebhookConfigurations().Informer(ctx)}, nil
	case admissionregistration.SchemeGroupVersion.WithResource("validatingwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Admissionregistration().InternalVersion().ValidatingWebhookConfigurations().Informer(ctx)}, nil

		// Group=apps, Version=internalVersion
	case apps.SchemeGroupVersion.WithResource("controllerrevisions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().InternalVersion().ControllerRevisions().Informer(ctx)}, nil
	case apps.SchemeGroupVersion.WithResource("statefulsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().InternalVersion().StatefulSets().Informer(ctx)}, nil

		// Group=autoscaling, Version=internalVersion
	case autoscaling.SchemeGroupVersion.WithResource("horizontalpodautoscalers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Autoscaling().InternalVersion().HorizontalPodAutoscalers().Informer(ctx)}, nil

		// Group=batch, Version=internalVersion
	case batch.SchemeGroupVersion.WithResource("cronjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().InternalVersion().CronJobs().Informer(ctx)}, nil
	case batch.SchemeGroupVersion.WithResource("jobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Batch().InternalVersion().Jobs().Informer(ctx)}, nil

		// Group=certificates.k8s.io, Version=internalVersion
	case certificates.SchemeGroupVersion.WithResource("certificatesigningrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certificates().InternalVersion().CertificateSigningRequests().Informer(ctx)}, nil

		// Group=coordination.k8s.io, Version=internalVersion
	case coordination.SchemeGroupVersion.WithResource("leases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Coordination().InternalVersion().Leases().Informer(ctx)}, nil

		// Group=core, Version=internalVersion
	case core.SchemeGroupVersion.WithResource("componentstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ComponentStatuses().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("configmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ConfigMaps().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("endpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Endpoints().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("events"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Events().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("limitranges"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().LimitRanges().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("namespaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Namespaces().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("nodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Nodes().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("persistentvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().PersistentVolumes().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("persistentvolumeclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().PersistentVolumeClaims().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("pods"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Pods().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("podtemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().PodTemplates().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("replicationcontrollers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ReplicationControllers().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ResourceQuotas().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("secrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Secrets().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("services"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().Services().Informer(ctx)}, nil
	case core.SchemeGroupVersion.WithResource("serviceaccounts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().InternalVersion().ServiceAccounts().Informer(ctx)}, nil

		// Group=extensions, Version=internalVersion
	case extensions.SchemeGroupVersion.WithResource("daemonsets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().DaemonSets().Informer(ctx)}, nil
	case extensions.SchemeGroupVersion.WithResource("deployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().Deployments().Informer(ctx)}, nil
	case extensions.SchemeGroupVersion.WithResource("ingresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().Ingresses().Informer(ctx)}, nil
	case extensions.SchemeGroupVersion.WithResource("replicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().InternalVersion().ReplicaSets().Informer(ctx)}, nil

		// Group=networking.k8s.io, Version=internalVersion
	case networking.SchemeGroupVersion.WithResource("networkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Networking().InternalVersion().NetworkPolicies().Informer(ctx)}, nil

		// Group=policy, Version=internalVersion
	case policy.SchemeGroupVersion.WithResource("poddisruptionbudgets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().InternalVersion().PodDisruptionBudgets().Informer(ctx)}, nil
	case policy.SchemeGroupVersion.WithResource("podsecuritypolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Policy().InternalVersion().PodSecurityPolicies().Informer(ctx)}, nil

		// Group=rbac.authorization.k8s.io, Version=internalVersion
	case rbac.SchemeGroupVersion.WithResource("clusterroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().ClusterRoles().Informer(ctx)}, nil
	case rbac.SchemeGroupVersion.WithResource("clusterrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().ClusterRoleBindings().Informer(ctx)}, nil
	case rbac.SchemeGroupVersion.WithResource("roles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().Roles().Informer(ctx)}, nil
	case rbac.SchemeGroupVersion.WithResource("rolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rbac().InternalVersion().RoleBindings().Informer(ctx)}, nil

		// Group=scheduling.k8s.io, Version=internalVersion
	case scheduling.SchemeGroupVersion.WithResource("priorityclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().InternalVersion().PriorityClasses().Informer(ctx)}, nil

		// Group=settings.k8s.io, Version=internalVersion
	case settings.SchemeGroupVersion.WithResource("podpresets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Settings().InternalVersion().PodPresets().Informer(ctx)}, nil

		// Group=storage.k8s.io, Version=internalVersion
	case storage.SchemeGroupVersion.WithResource("storageclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().InternalVersion().StorageClasses().Informer(ctx)}, nil
	case storage.SchemeGroupVersion.WithResource("volumeattachments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().InternalVersion().VolumeAttachments().Informer(ctx)}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
