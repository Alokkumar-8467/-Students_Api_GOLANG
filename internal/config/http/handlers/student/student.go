package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/alokMIPL/students-api/internal/types"
	"github.com/alokMIPL/students-api/internal/utils/response"
)
