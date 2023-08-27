package scenes

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/SkeleboyStudios/DigDeep/systems"
)

type MainMenuScene struct{}

func (*MainMenuScene) Type() string { return "Main Menu" }

func (s *MainMenuScene) Preload() {
	engo.Files.Load("scribish.ttf")
	engo.Files.Load("uipack_rpg_sheet.xml")
}

func (s *MainMenuScene) Setup(u engo.Updater) {
	w := u.(*ecs.World)

	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	var animatable *common.Animationable
	var notanimatable *common.NotAnimationable
	var animSys = &common.AnimationSystem{}
	w.AddSystemInterface(animSys, animatable, notanimatable)

	var audioable *common.Audioable
	var notaudioable *common.NotAudioable
	w.AddSystemInterface(&common.AudioSystem{}, audioable, notaudioable)

	var collisionable *common.Collisionable
	var notcollisionable *common.NotCollisionable
	w.AddSystemInterface(&common.CollisionSystem{}, collisionable, notcollisionable)

	var mouseable *common.Mouseable
	var notmouseable *common.NotMouseable
	w.AddSystemInterface(&common.MouseSystem{}, mouseable, notmouseable)

	var buttonable *systems.Buttonable
	var notbuttonable *systems.NotButtonable
	w.AddSystemInterface(&systems.ButtonSystem{}, buttonable, notbuttonable)

	//Title Text
	title := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
	}{BasicEntity: ecs.NewBasic()}
	titleFnt := &common.Font{
		URL:  "scribish.ttf",
		FG:   color.RGBA{0xB5, 0x94, 0x10, 0xFF},
		Size: 32,
	}
	titleFnt.CreatePreloaded()
	title.Drawable = common.Text{
		Font: titleFnt,
		Text: "Dig Deeper",
	}
	title.Position = engo.Point{X: 0, Y: 5}
	title.Scale = engo.Point{X: 3, Y: 3}
	title.SetZIndex(0)
	title.SetShader(common.TextHUDShader)
	w.AddEntity(&title)

	//Background
	bg := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
	}{BasicEntity: ecs.NewBasic()}
	bgTex, _ := common.LoadedSprite("panelInset_brown.png")
	bg.Drawable = bgTex
	bg.Position = engo.Point{X: 40, Y: 100}
	bg.Scale = engo.Point{X: 6, Y: 2.5}
	bg.SetZIndex(0)
	bg.SetShader(common.HUDShader)
	w.AddEntity(&bg)

	//New Game Button
	ngButton := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
		common.MouseComponent
		systems.ButtonComponent
	}{BasicEntity: ecs.NewBasic()}
	ngTex, _ := common.LoadedSprite("buttonLong_beige.png")
	ngButton.Drawable = ngTex
	ngButton.Position = engo.Point{X: 85, Y: 125}
	ngButton.Scale = engo.Point{X: 2.5, Y: 1}
	ngButton.Width = ngTex.Width() * ngButton.Scale.X
	ngButton.Height = ngTex.Height() * ngButton.Scale.Y
	ngButton.SetZIndex(1)
	ngButton.SetShader(common.HUDShader)
	ngButton.OnClick = func() {
		println("new game")
	}
	w.AddEntity(&ngButton)

	//New Game Text
	ngText := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
	}{BasicEntity: ecs.NewBasic()}
	ngText.Drawable = common.Text{
		Font: titleFnt,
		Text: "New Game!",
	}
	ngText.Position = engo.Point{X: 230, Y: 130}
	ngText.Scale = engo.Point{X: 1, Y: 1.25}
	ngText.SetZIndex(2)
	ngText.SetShader(common.TextHUDShader)
	w.AddEntity(&ngText)

	//Options Button
	optsButton := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
		common.MouseComponent
		systems.ButtonComponent
	}{BasicEntity: ecs.NewBasic()}
	optsTex, _ := common.LoadedSprite("buttonLong_blue.png")
	optsButton.Drawable = optsTex
	optsButton.Position = engo.Point{X: 85, Y: 200}
	optsButton.Scale = engo.Point{X: 2.5, Y: 1}
	optsButton.Width = optsTex.Width() * optsButton.Scale.X
	optsButton.Height = optsTex.Height() * optsButton.Scale.Y
	optsButton.SetZIndex(1)
	optsButton.SetShader(common.HUDShader)
	optsButton.OnClick = func() { println("Options") }
	w.AddEntity(&optsButton)

	//Options Text
	optsText := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
	}{BasicEntity: ecs.NewBasic()}
	optsText.Drawable = common.Text{
		Font: titleFnt,
		Text: "Options",
	}
	optsText.Position = engo.Point{X: 250, Y: 205}
	optsText.Scale = engo.Point{X: 1, Y: 1.25}
	optsText.SetZIndex(2)
	optsText.SetShader(common.TextHUDShader)
	w.AddEntity(&optsText)

	//Credits Button
	credButton := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
		common.MouseComponent
		systems.ButtonComponent
	}{BasicEntity: ecs.NewBasic()}
	credTex, _ := common.LoadedSprite("buttonLong_blue.png")
	credButton.Drawable = credTex
	credButton.Position = engo.Point{X: 85, Y: 275}
	credButton.Scale = engo.Point{X: 2.5, Y: 1}
	credButton.Width = credTex.Width() * credButton.Scale.X
	credButton.Height = credTex.Height() * credButton.Scale.Y
	credButton.SetZIndex(1)
	credButton.SetShader(common.HUDShader)
	credButton.OnClick = func() { println("Credits") }
	w.AddEntity(&credButton)

	//Credits Text
	credText := struct {
		ecs.BasicEntity
		common.SpaceComponent
		common.RenderComponent
	}{BasicEntity: ecs.NewBasic()}
	credText.Drawable = common.Text{
		Font: titleFnt,
		Text: "Credits",
	}
	credText.Position = engo.Point{X: 250, Y: 280}
	credText.Scale = engo.Point{X: 1, Y: 1.25}
	credText.SetZIndex(2)
	credText.SetShader(common.TextHUDShader)
	w.AddEntity(&credText)

}
