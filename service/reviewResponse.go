package main

import (
    v1 "k8s.io/api/admission/v1"

)


func allowAll(requestAdmissionReview v1.AdmissionReview) *v1.AdmissionResponse {
    reviewResponse := v1.AdmissionResponse{}
    reviewResponse.Allowed = true
    return &reviewResponse
}