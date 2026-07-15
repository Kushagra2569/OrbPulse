package orbpuls

import (
	"context"
	"fmt"
	"github.com/Kushagra2569/OrbPulse/internal/api"
	"net/http"
	"time"
)

func main() {
	fmt.Println("start")

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	ctx := context.Background()

	api.FetchCurrency(ctx, client)
}
