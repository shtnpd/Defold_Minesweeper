components {
  id: "tilemap"
  component: "/main/grid.tilemap"
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
  id: "game"
  component: "/main/game.script"
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
embedded_components {
  id: "point"
  type: "label"
  data: "size {\n  x: 128.0\n  y: 32.0\n  z: 0.0\n  w: 0.0\n}\ncolor {\n  x: 1.0\n  y: 1.0\n  z: 1.0\n  w: 1.0\n}\noutline {\n  x: 0.0\n  y: 0.0\n  z: 0.0\n  w: 1.0\n}\nshadow {\n  x: 0.0\n  y: 0.0\n  z: 0.0\n  w: 1.0\n}\nleading: 1.0\ntracking: 0.0\npivot: PIVOT_CENTER\nblend_mode: BLEND_MODE_ALPHA\nline_break: false\ntext: \"\"\nfont: \"/builtins/fonts/system_font.font\"\nmaterial: \"/builtins/fonts/label.material\"\n"
  position {
    x: 488.0
    y: 701.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 1.627
    y: 1.664
    z: 1.0
  }
}
