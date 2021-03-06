package stripe

import (
	"encoding/json"
	"testing"

	"github.com/fatsoma/tabb-stripe-go/form"
	assert "github.com/stretchr/testify/require"
)

func TestCardListParams_AppendTo(t *testing.T) {
	// Adds `object` for account (this will hit the external accounts endpoint)
	{
		params := &CardListParams{Account: String("acct_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"card"}, body.Get("object"))
	}

	// Adds `object` for customer (this will hit the sources endpoint)
	{
		params := &CardListParams{Customer: String("cus_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"card"}, body.Get("object"))
	}

	// *Doesn't* add `object` for recipient (this will hit the recipient cards
	// endpoint, so all possible resources are cards)
	{
		params := &CardListParams{Recipient: String("rp_123")}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string(nil), body.Get("object"))
	}
}

func TestCard_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Card
		err := json.Unmarshal([]byte(`"card_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Card{ID: "card_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "card_123", v.ID)
	}
}
