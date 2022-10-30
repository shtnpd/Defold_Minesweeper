components {
  id: "main"
  component: "/main/main.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
components {
  id: "start"
  component: "/main/start.label"
  position {
    x: -186.0
    y: 131.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
  scale {
    x: 1.315
    y: 1.73
    z: 1.0
  }
}
embedded_components {
  id: "background"
  type: "sprite"
  data: "tile_set: \"/main/main.atlas\"\ndefault_animation: \"checkboard\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\n"
  position {
    x: 2.0
    y: 137.0
    z: -1.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
