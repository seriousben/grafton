package acceptance

import (
	"context"
	"errors"
	"fmt"
	"time"

	gm "github.com/onsi/gomega"

	"github.com/manifoldco/go-manifold"
	"github.com/manifoldco/go-manifold/idtype"

	"github.com/manifoldco/grafton"
	"github.com/manifoldco/grafton/connector"
	"github.com/manifoldco/grafton/generated/provider/client/resource"
)

var errTimeout = errors.New("Exceeded Callback Wait time")
var resourceID manifold.ID
var curResource *connector.Resource

var provision = Feature("provision", "Provision a resource", func(ctx context.Context) {
	Default(func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		curResource, callbackID, async, err := provisionResource(ctx, api, product, plan, region)
		gm.Expect(err).To(gm.BeNil(), "Expected a successful provision of a resource")

		if async {
			c := fakeConnector.GetCallback(callbackID)

			gm.Expect(c.State).To(
				gm.Equal(connector.DoneCallbackState),
				"Expected to receive 'done' as the state",
			)
			gm.Expect(len(c.Message)).To(gm.SatisfyAll(
				gm.BeNumerically(">=", 3),
				gm.BeNumerically("<", 256),
			), "Message must be between 3 and 256 characters long.")
			gm.Expect(len(c.Credentials)).To(
				gm.Equal(0),
				"Credentials cannot be returned on a resource provisioning callback",
			)
		}

		resourceID = curResource.ID
		fakeConnector.AddResource(curResource)
	})

	ErrorCase("with a faulty product name", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		_, _, async, err := provisionResource(ctx, api, "", plan, region)

		gm.Expect(async).To(
			gm.BeFalse(),
			"Validation errors should be returned on the initial request",
		)
		gm.Expect(err).ShouldNot(
			gm.BeNil(),
			"Expected an error, got nil",
		)
		gm.Expect(err).Should(
			gm.BeAssignableToTypeOf(&resource.PutResourcesIDBadRequest{}),
			"Expected a BadRequest error, got %T", err,
		)
	})

	ErrorCase("with a faulty plan name", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		_, _, async, err := provisionResource(ctx, api, product, "faulty-plan-name", region)

		gm.Expect(async).To(
			gm.BeFalse(),
			"Validation errors should be returned on the initial request",
		)
		gm.Expect(err).ShouldNot(
			gm.BeNil(),
			"Expected an error, got nil",
		)
		gm.Expect(err).Should(
			gm.BeAssignableToTypeOf(&resource.PutResourcesIDBadRequest{}),
			"Expected a BadRequest error, got %T", err,
		)
	})

	ErrorCase("with a faulty region", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		_, _, async, err := provisionResource(ctx, api, product, plan, "faulty-region")

		gm.Expect(async).To(
			gm.BeFalse(),
			"Validation errors should be returned on the initial request",
		)
		gm.Expect(err).ShouldNot(
			gm.BeNil(),
			"Expected an error, got nil",
		)
		gm.Expect(err).Should(
			gm.BeAssignableToTypeOf(&resource.PutResourcesIDBadRequest{}),
			"Expected a BadRequest error, got %T", err,
		)
	})

	ErrorCase("with a bad signature", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		_, _, async, err := provisionResource(ctx, uapi, product, plan, region)

		gm.Expect(async).To(
			gm.BeFalse(),
			"Validation errors should be returned on the initial request",
		)
		gm.Expect(err).ShouldNot(
			gm.BeNil(),
			"Expected an error, got nil",
		)
		gm.Expect(err).Should(
			gm.BeAssignableToTypeOf(&resource.PutResourcesIDUnauthorized{}),
			"Expected an Unauthorized error, got %T", err,
		)
	})

	ErrorCase("with an already provisioned resource - same content acts as created", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		_, _, async, err := provisionResourceID(ctx, api, resourceID, product, plan, region)

		gm.Expect(async).To(
			gm.BeFalse(),
			"Same content should be evaluated during the initial call from Manifold",
		)
		gm.Expect(err).To(gm.BeNil(), "Create response should be returned (Repeatable Action)")
	})

	ErrorCase("with an already provisioned resource - different content results in conflict", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		var err error
		_, _, async, err := provisionResourceID(ctx, api, resourceID, product, newPlan, region)

		gm.Expect(async).To(
			gm.BeFalse(),
			"Same content should be evaluated during the initial call from Manifold",
		)
		gm.Expect(err).ShouldNot(
			gm.BeNil(),
			"Expected an error, got nil",
		)
		gm.Expect(err).Should(
			gm.BeAssignableToTypeOf(&resource.PutResourcesIDConflict{}),
			"Expected a 409 Conflict error, got %T (Repeated Action)", err,
		)
	})
})

var _ = provision.TearDown("Deprovision a resource", func(ctx context.Context) {
	Default(func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		callbackID, async, err := deprovisionResource(ctx, api, resourceID)

		gm.Expect(err).To(gm.BeNil(), "No error is expected")
		if async {
			c := fakeConnector.GetCallback(callbackID)

			gm.Expect(c.State).To(
				gm.Equal(connector.DoneCallbackState),
				"Expected to receive 'done' as the state",
			)
			gm.Expect(len(c.Message)).To(gm.SatisfyAll(
				gm.BeNumerically(">=", 3),
				gm.BeNumerically("<", 256),
			), "Message must be between 3 and 256 characters long.")
			gm.Expect(len(c.Credentials)).To(
				gm.Equal(0),
				"Credentials cannot be returned on a resource deprovisioning callback",
			)
		}
	})

	ErrorCase("delete a non existing resource", func() {
		ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()
		fakeID, _ := manifold.NewID(idtype.Resource)
		_, async, err := deprovisionResource(ctx, api, fakeID)

		gm.Expect(async).To(
			gm.BeFalse(),
			"Resource existence should be evaluated during the initial call from Manifold",
		)
		gm.Expect(err).ShouldNot(
			gm.BeNil(),
			"Expected an error, got nil",
		)
		gm.Expect(err).Should(
			gm.BeAssignableToTypeOf(&resource.DeleteResourcesIDNotFound{}),
			"Expected a resource NotFound error, got %T", err,
		)
	})
})
var _ = provision.RequiredFlags("product", "plan", "region", "new-plan")

func waitForCallback(ID manifold.ID, max time.Duration) (*connector.Callback, error) {
	timeout := time.After(max)

waitForCallback:
	select {
	case cb := <-fakeConnector.OnCallback:
		if cb.ID != ID {
			goto waitForCallback
		}

		if cb.State == connector.PendingCallbackState {
			goto waitForCallback
		}

		return cb, nil
	case <-timeout:
		return nil, errTimeout
	}
}

func provisionResource(ctx context.Context, api *grafton.Client, product, plan, region string) (*connector.Resource, manifold.ID, bool, error) {
	Infoln("Attempting to provision resource")

	ID, err := manifold.NewID(idtype.Resource)
	if err != nil {
		return nil, ID, false, FatalErr("Could not generate resource id:", err)
	}

	return provisionResourceID(ctx, api, ID, product, plan, region)
}

func provisionResourceID(ctx context.Context, api *grafton.Client, id manifold.ID, product, plan, region string) (*connector.Resource, manifold.ID, bool, error) {
	c, err := fakeConnector.AddCallback(connector.ResourceProvisionCallback)
	if err != nil {
		return nil, c.ID, false, err
	}

	msg, callback, err := api.ProvisionResource(ctx, c.ID, id, product, plan, region)
	if err != nil {
		return nil, c.ID, false, err
	}

	if callback {
		Infoln(fmt.Sprintf("Waiting for Callback (max: %.1f minutes): %s", cbTimeout.Minutes(), msg))

		cb, err := waitForCallback(c.ID, cbTimeout)
		if err != nil {
			return nil, c.ID, callback, err
		}

		msg = cb.Message
	}

	Infoln("Resource Provisioned Successfully:", id)
	if msg != "" {
		Infoln("Message: ", msg)
	}

	return &connector.Resource{
		ID:        id,
		Product:   product,
		Plan:      plan,
		Region:    region,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, c.ID, callback, nil
}

func deprovisionResource(ctx context.Context, api *grafton.Client, resourceID manifold.ID) (manifold.ID, bool, error) {
	Infoln("Attempting to deprovision resource:", resourceID)

	c, err := fakeConnector.AddCallback(connector.ResourceDeprovisionCallback)
	if err != nil {
		return manifold.ID{}, false, err
	}

	msg, callback, err := api.DeprovisionResource(ctx, c.ID, resourceID)
	if err != nil {
		return c.ID, callback, err
	}

	if callback {
		Infoln(fmt.Sprintf("Waiting for Callback (max: %.1f minutes): %s", cbTimeout.Minutes(), msg))

		cb, err := waitForCallback(c.ID, cbTimeout)
		if err != nil {
			return c.ID, callback, err
		}

		msg = cb.Message
	}

	Infoln("Resource Deprovisioned.")
	if msg != "" {
		Infoln("Callback Message: ", msg)
	}

	return c.ID, callback, nil
}