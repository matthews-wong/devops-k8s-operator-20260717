package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	operatorsv1 "github.com/matthews-wong/devops-k8s-operator/api/v1"
)

// PingReconciler reconciles a Ping object.
type PingReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=operators.example.com,resources=pings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operators.example.com,resources=pings/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=operators.example.com,resources=pings/finalizers,verbs=update

// Reconcile is the main loop for the Ping controller.
func (r *PingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var ping operatorsv1.Ping
	if err := r.Get(ctx, req.NamespacedName, &ping); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	interval := ping.Spec.IntervalSeconds
	if interval <= 0 {
		interval = 30
	}
	msg := ping.Spec.Message
	if msg == "" {
		msg = "ping"
	}

	// Emit a ping on the log; idempotent within the same second.
	now := time.Now().UTC().Format(time.RFC3339)
	logger.Info("ping", "message", msg, "name", req.Name)

	if ping.Status.LastPinged != now {
		ping.Status.LastPinged = now
		ping.Status.PingCount++
		if err := r.Status().Update(ctx, &ping); err != nil {
			logger.Error(err, "failed to update status")
			return ctrl.Result{RequeueAfter: time.Duration(interval) * time.Second}, err
		}
	}

	return ctrl.Result{RequeueAfter: time.Duration(interval) * time.Second}, nil
}

// SetupWithManager wires the reconciler into the manager.
func (r *PingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorsv1.Ping{}).
		Complete(r)
}
