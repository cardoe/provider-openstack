/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-openstack/apis/networking/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ClusterV1.
func (mg *ClusterV1) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FixedNetwork),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FixedNetworkRef,
		Selector:     mg.Spec.ForProvider.FixedNetworkSelector,
		To: reference.To{
			List:    &v1alpha1.NetworkV2List{},
			Managed: &v1alpha1.NetworkV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FixedNetwork")
	}
	mg.Spec.ForProvider.FixedNetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FixedNetworkRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FixedSubnet),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FixedSubnetRef,
		Selector:     mg.Spec.ForProvider.FixedSubnetSelector,
		To: reference.To{
			List:    &v1alpha1.SubnetV2List{},
			Managed: &v1alpha1.SubnetV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FixedSubnet")
	}
	mg.Spec.ForProvider.FixedSubnet = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FixedSubnetRef = rsp.ResolvedReference

	return nil
}
