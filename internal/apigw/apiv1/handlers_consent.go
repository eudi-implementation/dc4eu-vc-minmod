package apiv1

import (
	"context"
	"time"
	"vc/internal/apigw/db"
	"vc/pkg/model"
)

// AddConsentRequest is the request for AddConsent
type AddConsentRequest struct {
	AuthenticSource         string `json:"authentic_source" validate:"required"`
	AuthenticSourcePersonID string `json:"authentic_source_person_id" validate:"required"`
	ConsentTo               string `json:"consent_to"`
	SessionID               string `json:"session_id"`
}

// AddConsent adds a consent to a document
//
//	@Summary		AddConsent
//	@ID				add-consent
//	@Description	Add consent endpoint
//	@Tags			dc4eu
//	@Accept			json
//	@Produce		json
//	@Success		200	"Success"
//	@Failure		400	{object}	helpers.ErrorResponse	"Bad Request"
//	@Param			req	body		AddConsentRequest		true	" "
//	@Router			/consent [post]
func (c *Client) AddConsent(ctx context.Context, req *AddConsentRequest) error {
	err := c.db.VCConsentColl.Add(ctx, &db.AddConsentQuery{
		AuthenticSource:         req.AuthenticSource,
		AuthenticSourcePersonID: req.AuthenticSourcePersonID,
		Consent: &model.Consent{
			ConsentTo: req.ConsentTo,
			SessionID: req.SessionID,
			CreatedAt: time.Now().Unix(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// GetConsentRequest is the request for GetConsent
type GetConsentRequest struct {
	AuthenticSource         string `json:"authentic_source" validate:"required"`
	AuthenticSourcePersonID string `json:"authentic_source_person_id" validate:"required"`
}

// GetConsent gets a consent for a document
//
//	@Summary		GetConsent
//	@ID				get-consent
//	@Description	Get consent endpoint
//	@Tags			dc4eu
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} model.Consent	"Success"
//	@Failure		400	{object}		helpers.ErrorResponse	"Bad Request"
//	@Param			req	body			GetConsentRequest		true	" "
//	@Router			/consent/get [post]
func (c *Client) GetConsent(ctx context.Context, req *GetConsentRequest) (*model.Consent, error) {
	res, err := c.db.VCConsentColl.Get(ctx, &db.GetConsentQuery{
		AuthenticSource:         req.AuthenticSource,
		AuthenticSourcePersonID: req.AuthenticSourcePersonID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
