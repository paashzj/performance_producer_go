package conf

import (
	"github.com/paashzj/gutil"
)

var (
	RoutineNum = gutil.GetEnvInt("ROUTINE_NUM", 1)
)
