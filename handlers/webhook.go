package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/alertmanager/template"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := template.Data{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		asJson(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Printf("Alerts: GroupLabels=%v, CommonLabels=%v, Receiver=%v", data.GroupLabels, data.CommonLabels, data.Receiver)

	if data.Receiver != "nodeDownReceiver" {
		asJson(w, http.StatusNotModified, "Alert not meant for handler")
		return
	}
	// creates the in-cluster kube config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	for _, alert := range data.Alerts {
		log.Printf("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)

		nodeName := alert.Annotations["nodeName"]
		if err := clientset.CoreV1().Nodes().Delete(ctx, nodeName, metav1.DeleteOptions{}); err != nil {
			asJson(w, http.StatusInternalServerError, "Failed deleting node")
		}
	}

	asJson(w, http.StatusOK, "success")
}
