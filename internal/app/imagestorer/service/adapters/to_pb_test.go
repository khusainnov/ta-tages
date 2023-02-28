package adapters

import (
	"testing"
	"time"

	"github.com/khusainnov/tag/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImageToPb(t *testing.T) {
	t.Run("ok", func(t *testing.T) {

	})
}

func TestListImageToPb(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		list := []*model.Image{
			{
				Name:      "eab6afa7-ce1f-4f7d-b348-ae96bef41aaa",
				CreatedAt: time.Date(2023, time.February, 28, 22, 13, 0, 0, time.UTC),
				EditedAt:  time.Date(2023, time.February, 28, 22, 15, 0, 0, time.UTC),
			},
			{
				Name:      "02486af9-84df-4238-99fe-49f21a53644b",
				CreatedAt: time.Date(2021, time.September, 2, 11, 13, 43, 0, time.UTC),
				EditedAt:  time.Date(2022, time.December, 16, 23, 17, 53, 0, time.UTC),
			},
			{
				Name:      "fc6f1b67-bd3a-4700-9566-a6939ee24688",
				CreatedAt: time.Date(2015, time.November, 6, 6, 6, 10, 0, time.UTC),
				EditedAt:  time.Date(2020, time.May, 9, 6, 23, 32, 0, time.UTC),
			},
		}

		expected := []struct {
			Name      string
			CreatedAt *timestamppb.Timestamp
			EditedAt  *timestamppb.Timestamp
			Err       error
		}{
			{
				Name: "eab6afa7-ce1f-4f7d-b348-ae96bef41aaa",
				CreatedAt: &timestamppb.Timestamp{
					Seconds: 1677622380,
					Nanos:   0,
				},
				EditedAt: &timestamppb.Timestamp{
					Seconds: 1677622500,
					Nanos:   0,
				},
				Err: nil,
			},
			{
				Name: "02486af9-84df-4238-99fe-49f21a53644b",
				CreatedAt: &timestamppb.Timestamp{
					Seconds: 1630581223,
					Nanos:   0,
				},
				EditedAt: &timestamppb.Timestamp{
					Seconds: 1671232673,
					Nanos:   0,
				},
				Err: nil,
			},
			{
				Name: "fc6f1b67-bd3a-4700-9566-a6939ee24688",
				CreatedAt: &timestamppb.Timestamp{
					Seconds: 1446789970,
					Nanos:   0,
				},
				EditedAt: &timestamppb.Timestamp{
					Seconds: 1589005412,
					Nanos:   0,
				},
				Err: nil,
			},
		}

		resp := ListImageToPb(list)

		for i, v := range resp.Images {
			assert.Equal(t, expected[i].Name, v.Name)
			require.Equal(t, expected[i].CreatedAt, v.CreatedDate)
			require.Equal(t, expected[i].EditedAt, v.ModifiedDate)
		}
	})
}
