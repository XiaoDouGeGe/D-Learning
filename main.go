package main

import (
	"dl/config"
	"dl/controller"
	"dl/database"
	"dl/service"

	"dl/utils"
	"fmt"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := newApp()
	//应用App设置
	configuration(app)
	//路由设置
	mvcHandle(app)
	config := config.InitConfig()
	fmt.Println(config)
	addr := ":" + config.ServerPort
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
}

// 构建App
func newApp() *iris.Application {
	app := iris.New()
	// app.Use(irisMiddleware)
	app.Use(Cors)

	// 设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")
	//注册静态资源 获取外部配置的静态资源配置位置
	config := config.InitConfig()
	resourceDir := config.ResourceDir
	app.HandleDir("/api/res/", resourceDir)
	return app
}

func irisMiddleware(ctx iris.Context) {
	utils.LoggerInfo(ctx.Path())
	ctx.Next()
}

// Cors
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

/**
 * MVC 架构模式处理
 */
func mvcHandle(app *iris.Application) {
	//实例化数据库引擎
	engine := database.NewEngine()

	userService := service.NewUserService(engine)
	courseService := service.NewCourseService(engine)
	chapterService := service.NewChapterService(engine)
	progressService := service.NewProgressService(engine)
	chatService := service.NewChatService(engine)
	questionService := service.NewQuestionService(engine)
	choiceService := service.NewChoiceService(engine)
	exerciseService := service.NewExerciseService(engine)

	user := mvc.New(app.Party("/api/user"))
	user.Register(userService)
	user.Handle(new(controller.UserController))

	course := mvc.New(app.Party("/api/course"))
	course.Register(courseService, userService, chapterService, progressService)
	course.Handle(new(controller.CourseController))

	chapter := mvc.New(app.Party("/api/chapter"))
	chapter.Register(chapterService)
	chapter.Handle(new(controller.ChapterController))

	progress := mvc.New(app.Party("/api/progress"))
	progress.Register(progressService)
	progress.Handle(new(controller.ProgressController))

	chat := mvc.New(app.Party("/api/chat"))
	chat.Register(chatService)
	chat.Handle(new(controller.ChatController))

	question := mvc.New(app.Party("/api/question"))
	question.Register(questionService, choiceService)
	question.Handle(new(controller.QuestionController))

	exercise := mvc.New(app.Party("/api/exercise"))
	exercise.Register(exerciseService, chapterService, userService)
	exercise.Handle(new(controller.ExerciseController))

}

/**
 * 项目设置
 */
func configuration(app *iris.Application) {
	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))
	//错误配置
	app.OnErrorCode(iris.StatusNotFound, func(context *context.Context) {
		context.JSON(iris.Map{
			"error_msg_en": "404 not found",
			"error_msg_zh": "404 not found",
		})
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context *context.Context) {
		context.JSON(iris.Map{
			"error_msg_en": "500 internal error",
			"error_msg_zh": "500 internal error",
		})
	})
}
