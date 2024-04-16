package endpoints

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"pointsCounterCRUD/database"
	"pointsCounterCRUD/endpoints/controller"
	"strconv"
)

type Routs struct {
	db   *gorm.DB
	r    *gin.Engine
	repo *database.Repo
}

func NewRouts(repo *database.Repo) *Routs {
	if repo == nil {
		//because i can
		log.Fatalln("Database needed!!")
	}
	return &Routs{
		repo: repo,
		db:   repo.DB,
		r:    gin.Default(),
	}
}
func (r *Routs) AddAuthPaths() {

	authRoutes := r.r.Group("/auth/user")
	{
		// registration route
		authRoutes.POST("/register", controller.Register)
		// login route
		authRoutes.POST("/login", controller.Login)

	}
}

func (r *Routs) AddPaths() {
	//base := NewBasicForRouts(r.db)

	r.r.GET("/", halfCheck)
	r.r.GET("/tea", coffee)
	r.r.GET("/coffee", coffee)
}

func (r *Routs) Start(port int16) {
	portString := ":" + strconv.Itoa(int(port))
	if err := r.r.Run(portString); err != nil {
		log.Fatalln("gin do not start on", port)
	}
}
