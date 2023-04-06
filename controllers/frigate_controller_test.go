package controllers

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	// "k8s.io/client-go/testing"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	// api "github.com/example/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	webappv1 "my.domain/api/v1"
)

func TestMyCustomResourceReconciler(t *testing.T) {
	// Create a fake clientset.
	scheme := runtime.NewScheme()
	err := webappv1.AddToScheme(scheme)
	assert.NoError(t, err)

	// Create a new instance of your CRD.
	myCRD := &webappv1.Frigate{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-crd",
			Namespace: "default",
		},
		Spec: webappv1.FrigateSpec{
			Foo:       "bar",
			FirstName: "Ningan",
			LastName:  "An",
		},
	}

	// client := fake.NewSimpleClientset()
	client := fake.NewClientBuilder().WithScheme(scheme).WithObjects(myCRD).Build()

	// Create a new controller object with the fake clientset.
	controller := &FrigateReconciler{
		Client: client,
		Scheme: scheme,
	}

	// Test creating a new CRD.
	t.Run("TestCreateFrigate", func(t *testing.T) {
		controller.Reconcile(context.Background(), reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      myCRD.Name,
				Namespace: myCRD.Namespace,
			},
		})

		fmt.Printf("111: %+v\n", myCRD)
		// Ensure the CRD was created.
		err = client.Get(context.Background(), types.NamespacedName{Name: myCRD.Name, Namespace: myCRD.Namespace}, myCRD)

		fmt.Printf("222: %+v\n", myCRD)

		assert.NoError(t, err)
		// assert.Equal(t, myCRD.Spec.Foo, created.Spec.Foo)
	})

	// Test updating an existing CRD.
	t.Run("TestUpdateMyCustomResource", func(t *testing.T) {
		// Update the CRD.

		fmt.Printf("333: %+v\n", myCRD)
		myCRD.Spec.Foo = "baz"
		client.Update(context.Background(), myCRD)
		assert.NoError(t, err)

		controller.Reconcile(context.Background(), reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      myCRD.Name,
				Namespace: myCRD.Namespace,
			},
		})

		// Ensure the CRD was updated.
		err = client.Get(context.Background(), types.NamespacedName{Name: myCRD.Name, Namespace: myCRD.Namespace}, myCRD)

		fmt.Printf("444: %+v\n", myCRD)
		assert.NoError(t, err)
		// assert.Equal(t, myCRD.Spec.Foo, updated.Spec.Foo)
	})

	// Test deleting an existing CRD.
	t.Run("TestDeleteMyCustomResource", func(t *testing.T) {
		// Delete the CRD.
		err := client.Delete(context.Background(), myCRD)

		assert.NoError(t, err)

		controller.Reconcile(context.Background(), reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      myCRD.Name,
				Namespace: myCRD.Namespace,
			},
		})

		// Ensure the CRD was deleted.
		err = client.Get(context.Background(), types.NamespacedName{Name: myCRD.Name, Namespace: myCRD.Namespace}, myCRD)
		assert.Error(t, err)
	})
}
