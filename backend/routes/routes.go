package routes

import (
	handler "OJ-backend/controllers"
	"OJ-backend/services/sse"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Public routes
	e.POST("/login", handler.Login)
	e.POST("/admin/login", handler.AdminLogin)
	e.GET("/contests", handler.GetAllContests)

	// Callback route for workers (HMAC authenticated)
	e.POST("/callback/submission", handler.HandleSubmissionCallback)

	// Protected routes
	api := e.Group("/api")
	api.Use(handler.JWTMiddleware())
	api.GET("/profile", handler.GetProfile)
	api.PUT("/profile", handler.UpdateProfile)
	api.GET("/problems/:id", handler.GetAllProblemsByContestID)
	api.GET("/problem/:id", handler.GetProblemByID)
	api.GET("/testcases/:id", handler.GetAllTestCasesByProblemID)
	api.POST("/submit/:user_id/:problem_id", handler.HandleSubmission)
	api.GET("/leaderboard/:contest_id", handler.GetLeaderboardByContestID)
	// SSE endpoint for real-time submission updates
	e.GET("/submission/:user_id/:submission_id/events", sse.HandleSSEConnection)

	// Admin routes
	admin := e.Group("/admin")
	admin.Use(handler.AdminJWTMiddleware())
	//contest routes
	admin.POST("/create-contest", handler.CreateContest)
	admin.PUT("/contest/:id", handler.UpdateContest)
	admin.DELETE("/contest/:id", handler.DeleteContest)
	//problem routes
	admin.POST("/create-problem/:id", handler.CreateProblem)
	admin.GET("/problems/:id", handler.GetAllProblemsByContestID)
	admin.PUT("/problem/:id", handler.UpdateProblem)
	admin.DELETE("/problem/:id", handler.DeleteProblem)
	//test case routes
	admin.POST("/create-testcase/:id", handler.CreateTestCase)
	admin.GET("/testcases/:id", handler.GetAllTestCasesByProblemID)
	admin.PUT("/testcase/:id", handler.UpdateTestCase)
	admin.DELETE("/testcase/:id", handler.DeleteTestCase)
}
