/*
Copyright 2023 The Kubernetes authors.

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

package controllers

import (
	"context"
	"strings"
	"time"

	"golang.org/x/exp/slog"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	webappv1 "my.domain/api/v1"
)

// FrigateReconciler reconciles a Frigate object
type FrigateReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.example.com,resources=frigates,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.example.com,resources=frigates/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.example.com,resources=frigates/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Frigate object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *FrigateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	var result ctrl.Result
	obj := webappv1.Frigate{}
	// retryCount := 10
RETRY:
	err := r.Get(ctx, req.NamespacedName, &obj)
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			// retryCount--
			// if retryCount >= 0 {
			// 	goto RETRY
			// }
		} else {
			// if !r.notFound(req.NamespacedName, err) {
			result.Requeue = true
			result.RequeueAfter = time.Millisecond * 100
			return result, err
		}
	} else if r.matched(&obj) {
		// r.NotFound.Delete(req.Namespace + "/" + req.Name)
		if obj.DeletionTimestamp.IsZero() {
			// DeletionTimestamp 为空时,为创建或更新事件
			err := r.createOrUpdateCR(ctx, &obj)
			if err != nil {
				if strings.Contains(err.Error(), "Operation cannot be fulfilled") {
					// time.Sleep(time.Millisecond * 100)
					goto RETRY
				} else {
					result.Requeue = true
					result.RequeueAfter = time.Second
				}
				return result, err
			}
		} else {
			// DeletionTimestamp 不为空时,为删除事件
			err = r.deleteCR(ctx, &obj)
			if err != nil {
				if strings.Contains(err.Error(), "Operation cannot be fulfilled") {
					// time.Sleep(time.Millisecond * 100)
					goto RETRY
				} else {
					result.Requeue = true
					result.RequeueAfter = time.Second
				}
				return result, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FrigateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Frigate{}).
		Complete(r)
}

// matched 检查是否需要处理
func (r *FrigateReconciler) matched(obj *webappv1.Frigate) bool {

	return true
}

// createOrUpdate 检查是否需要处理
func (r *FrigateReconciler) createOrUpdateCR(ctx context.Context, obj *webappv1.Frigate) error {
	iNodeKey := obj.Namespace + "/" + obj.Name
	slog.Info("CreateOrUpdate frigate", "obj", iNodeKey)

		// 检查并填充finalizers
		fl := false
		for _, finalizer := range obj.GetFinalizers() {
			if finalizer == "123456789" {
				fl = true
				break
			}
		}
		if !fl {
			obj.Finalizers = append(obj.Finalizers, "123456789")
			err := r.Update(ctx, obj)
			if err != nil {
				return err
			}
			return nil
		}

	return nil
}

// delete 检查是否需要处理
func (r *FrigateReconciler) deleteCR(ctx context.Context, obj *webappv1.Frigate) error {
	iNodeKey := obj.Namespace + "/" + obj.Name
	slog.Info("Delete frigate", "obj", iNodeKey)

	obj.SetFinalizers([]string{})
	err := r.Update(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}
