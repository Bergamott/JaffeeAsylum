components {
  id: "walker"
  component: "/main/walker.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"tall_walk\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/mobiles.atlas\"\n"
  "}\n"
  ""
}
