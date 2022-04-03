 

 // Code generated by cmd/blockhash; DO NOT EDIT. 

  

 package block 

  

 const ( 

         hashAir = iota 

         hashAmethystBlock 

         hashAncientDebris 

         hashAndesite 

         hashBarrel 

         hashBarrier 

         hashBasalt 

         hashBeacon 

         hashBedrock 

         hashBeetrootSeeds 

         hashBlueIce 

         hashBoneBlock 

         hashBookshelf 

         hashBricks 

         hashCake 

         hashCalcite 

         hashCarpet 

         hashCarrot 

         hashChain 

         hashChest 

         hashChiseledQuartz 

         hashClay 

         hashCoalBlock 

         hashCoalOre 

         hashCobblestone 

         hashCocoaBean 

         hashConcrete 

         hashConcretePowder 

         hashCopperOre 

         hashCoral 

         hashCoralBlock 

         hashDeadBush 

         hashDiamondBlock 

         hashDiamondOre 

         hashDiorite 

         hashDirt 

         hashDirtPath 

         hashDoubleFlower 

         hashDoubleTallGrass 

         hashDragonEgg 

         hashDriedKelpBlock 

         hashDripstone 

         hashEmeraldBlock 

         hashEmeraldOre 

         hashEndBrickStairs 

         hashEndBricks 

         hashEndStone 

         hashFarmland 

         hashFire 

         hashFlower 

         hashGildedBlackstone 

         hashGlass 

         hashGlassPane 

         hashGlazedTerracotta 

         hashGlowstone 

         hashGoldBlock 

         hashGoldOre 

         hashGranite 

         hashGrass 

         hashGravel 

         hashHoneycombBlock 

         hashInvisibleBedrock 

         hashIronBars 

         hashIronBlock 

         hashIronOre 

         hashItemFrame 

         hashKelp 

         hashLadder 

         hashLantern 

         hashLapisBlock 

         hashLapisOre 

         hashLava 

         hashLeaves 

         hashLight 

         hashLitPumpkin 

         hashLog 

         hashMelon 

         hashMelonSeeds 

         hashMossCarpet 

         hashNetherBrickFence 

         hashNetherBricks 

         hashNetherGoldOre 

         hashNetherQuartzOre 

         hashNetherSprouts 

         hashNetherWart 

         hashNetheriteBlock 

         hashNetherrack 

         hashNoteBlock 

         hashObsidian 

         hashPackedIce 

         hashPlanks 

         hashPodzol 

         hashPotato 

         hashPrismarine 

         hashPumpkin 

         hashPumpkinSeeds 

         hashQuartz 

         hashQuartzBricks 

         hashQuartzPillar 

         hashRawCopperBlock 

         hashRawGoldBlock 

         hashRawIronBlock 

         hashSand 

         hashSandstone 

         hashSandstoneStairs 

         hashSeaLantern 

         hashSeaPickle 

         hashShroomlight 

         hashSign 

         hashSnow 

         hashSoulSand 

         hashSoulSoil 

         hashSponge 

         hashSporeBlossom 

         hashStainedGlass 

         hashStainedGlassPane 

         hashStainedTerracotta 

         hashStone 

         hashStoneBricks 

         hashTallGrass 

         hashTerracotta 

         hashTorch 

         hashTuff 

         hashWater 

         hashWheatSeeds 

         hashWood 

         hashWoodDoor 

         hashWoodFence 

         hashWoodFenceGate 

         hashWoodSlab 

         hashWoodStairs 

         hashWoodTrapdoor 

         hashWool 

 ) 

  

 func (Air) Hash() uint64 { 

         return hashAir 

 } 

  

 func (AmethystBlock) Hash() uint64 { 

         return hashAmethystBlock 

 } 

  

 func (AncientDebris) Hash() uint64 { 

         return hashAncientDebris 

 } 

  

 func (a Andesite) Hash() uint64 { 

         return hashAndesite | uint64(boolByte(a.Polished))<<8 

 } 

  

 func (b Barrel) Hash() uint64 { 

         return hashBarrel | uint64(b.Facing)<<8 | uint64(boolByte(b.Open))<<11 

 } 

  

 func (Barrier) Hash() uint64 { 

         return hashBarrier 

 } 

  

 func (b Basalt) Hash() uint64 { 

         return hashBasalt | uint64(boolByte(b.Polished))<<8 | uint64(b.Axis)<<9 

 } 

  

 func (Beacon) Hash() uint64 { 

         return hashBeacon 

 } 

  

 func (b Bedrock) Hash() uint64 { 

         return hashBedrock | uint64(boolByte(b.InfiniteBurning))<<8 

 } 

  

 func (b BeetrootSeeds) Hash() uint64 { 

         return hashBeetrootSeeds | uint64(b.Growth)<<8 

 } 

  

 func (BlueIce) Hash() uint64 { 

         return hashBlueIce 

 } 

  

 func (b BoneBlock) Hash() uint64 { 

         return hashBoneBlock | uint64(b.Axis)<<8 

 } 

  

 func (Bookshelf) Hash() uint64 { 

         return hashBookshelf 

 } 

  

 func (Bricks) Hash() uint64 { 

         return hashBricks 

 } 

  

 func (c Cake) Hash() uint64 { 

         return hashCake | uint64(c.Bites)<<8 

 } 

  

 func (Calcite) Hash() uint64 { 

         return hashCalcite 

 } 

  

 func (c Carpet) Hash() uint64 { 

         return hashCarpet | uint64(c.Colour.Uint8())<<8 

 } 

  

 func (c Carrot) Hash() uint64 { 

         return hashCarrot | uint64(c.Growth)<<8 

 } 

  

 func (c Chain) Hash() uint64 { 

         return hashChain | uint64(c.Axis)<<8 

 } 

  

 func (c Chest) Hash() uint64 { 

         return hashChest | uint64(c.Facing)<<8 

 } 

  

 func (ChiseledQuartz) Hash() uint64 { 

         return hashChiseledQuartz 

 } 

  

 func (Clay) Hash() uint64 { 

         return hashClay 

 } 

  

 func (CoalBlock) Hash() uint64 { 

         return hashCoalBlock 

 } 

  

 func (c CoalOre) Hash() uint64 { 

         return hashCoalOre | uint64(c.Type.Uint8())<<8 

 } 

  

 func (c Cobblestone) Hash() uint64 { 

         return hashCobblestone | uint64(boolByte(c.Mossy))<<8 

 } 

  

 func (c CocoaBean) Hash() uint64 { 

         return hashCocoaBean | uint64(c.Facing)<<8 | uint64(c.Age)<<10 

 } 

  

 func (c Concrete) Hash() uint64 { 

         return hashConcrete | uint64(c.Colour.Uint8())<<8 

 } 

  

 func (c ConcretePowder) Hash() uint64 { 

         return hashConcretePowder | uint64(c.Colour.Uint8())<<8 

 } 

  

 func (c CopperOre) Hash() uint64 { 

         return hashCopperOre | uint64(c.Type.Uint8())<<8 

 } 

  

 func (c Coral) Hash() uint64 { 

         return hashCoral | uint64(c.Type.Uint8())<<8 | uint64(boolByte(c.Dead))<<11 

 } 

  

 func (c CoralBlock) Hash() uint64 { 

         return hashCoralBlock | uint64(c.Type.Uint8())<<8 | uint64(boolByte(c.Dead))<<11 

 } 

  

 func (DeadBush) Hash() uint64 { 

         return hashDeadBush 

 } 

  

 func (DiamondBlock) Hash() uint64 { 

         return hashDiamondBlock 

 } 

  

 func (d DiamondOre) Hash() uint64 { 

         return hashDiamondOre | uint64(d.Type.Uint8())<<8 

 } 

  

 func (d Diorite) Hash() uint64 { 

         return hashDiorite | uint64(boolByte(d.Polished))<<8 

 } 

  

 func (d Dirt) Hash() uint64 { 

         return hashDirt | uint64(boolByte(d.Coarse))<<8 

 } 

  

 func (DirtPath) Hash() uint64 { 

         return hashDirtPath 

 } 

  

 func (d DoubleFlower) Hash() uint64 { 

         return hashDoubleFlower | uint64(boolByte(d.UpperPart))<<8 | uint64(d.Type.Uint8())<<9 

 } 

  

 func (d DoubleTallGrass) Hash() uint64 { 

         return hashDoubleTallGrass | uint64(boolByte(d.UpperPart))<<8 | uint64(d.Type.Uint8())<<9 

 } 

  

 func (DragonEgg) Hash() uint64 { 

         return hashDragonEgg 

 } 

  

 func (DriedKelpBlock) Hash() uint64 { 

         return hashDriedKelpBlock 

 } 

  

 func (Dripstone) Hash() uint64 { 

         return hashDripstone 

 } 

  

 func (EmeraldBlock) Hash() uint64 { 

         return hashEmeraldBlock 

 } 

  

 func (e EmeraldOre) Hash() uint64 { 

         return hashEmeraldOre | uint64(e.Type.Uint8())<<8 

 } 

  

 func (s EndBrickStairs) Hash() uint64 { 

         return hashEndBrickStairs | uint64(boolByte(s.UpsideDown))<<8 | uint64(s.Facing)<<9 

 } 

  

 func (EndBricks) Hash() uint64 { 

         return hashEndBricks 

 } 

  

 func (EndStone) Hash() uint64 { 

         return hashEndStone 

 } 

  

 func (f Farmland) Hash() uint64 { 

         return hashFarmland | uint64(f.Hydration)<<8 

 } 

  

 func (f Fire) Hash() uint64 { 

         return hashFire | uint64(f.Type.Uint8())<<8 | uint64(f.Age)<<9 

 } 

  

 func (f Flower) Hash() uint64 { 

         return hashFlower | uint64(f.Type.Uint8())<<8 

 } 

  

 func (GildedBlackstone) Hash() uint64 { 

         return hashGildedBlackstone 

 } 

  

 func (Glass) Hash() uint64 { 

         return hashGlass 

 } 

  

 func (GlassPane) Hash() uint64 { 

         return hashGlassPane 

 } 

  

 func (t GlazedTerracotta) Hash() uint64 { 

         return hashGlazedTerracotta | uint64(t.Colour.Uint8())<<8 | uint64(t.Facing)<<12 

 } 

  

 func (Glowstone) Hash() uint64 { 

         return hashGlowstone 

 } 

  

 func (GoldBlock) Hash() uint64 { 

         return hashGoldBlock 

 } 

  

 func (g GoldOre) Hash() uint64 { 

         return hashGoldOre | uint64(g.Type.Uint8())<<8 

 } 

  

 func (g Granite) Hash() uint64 { 

         return hashGranite | uint64(boolByte(g.Polished))<<8 

 } 

  

 func (Grass) Hash() uint64 { 

         return hashGrass 

 } 

  

 func (Gravel) Hash() uint64 { 

         return hashGravel 

 } 

  

 func (HoneycombBlock) Hash() uint64 { 

         return hashHoneycombBlock 

 } 

  

 func (InvisibleBedrock) Hash() uint64 { 

         return hashInvisibleBedrock 

 } 

  

 func (IronBars) Hash() uint64 { 

         return hashIronBars 

 } 

  

 func (IronBlock) Hash() uint64 { 

         return hashIronBlock 

 } 

  

 func (i IronOre) Hash() uint64 { 

         return hashIronOre | uint64(i.Type.Uint8())<<8 

 } 

  

 func (i ItemFrame) Hash() uint64 { 

         return hashItemFrame | uint64(i.Facing)<<8 | uint64(boolByte(i.Glowing))<<11 

 } 

  

 func (k Kelp) Hash() uint64 { 

         return hashKelp | uint64(k.Age)<<8 

 } 

  

 func (l Ladder) Hash() uint64 { 

         return hashLadder | uint64(l.Facing)<<8 

 } 

  

 func (l Lantern) Hash() uint64 { 

         return hashLantern | uint64(boolByte(l.Hanging))<<8 | uint64(l.Type.Uint8())<<9 

 } 

  

 func (LapisBlock) Hash() uint64 { 

         return hashLapisBlock 

 } 

  

 func (l LapisOre) Hash() uint64 { 

         return hashLapisOre | uint64(l.Type.Uint8())<<8 

 } 

  

 func (l Lava) Hash() uint64 { 

         return hashLava | uint64(boolByte(l.Still))<<8 | uint64(l.Depth)<<9 | uint64(boolByte(l.Falling))<<17 

 } 

  

 func (l Leaves) Hash() uint64 { 

         return hashLeaves | uint64(l.Wood.Uint8())<<8 | uint64(boolByte(l.Persistent))<<11 | uint64(boolByte(l.ShouldUpdate))<<12 

 } 

  

 func (l Light) Hash() uint64 { 

         return hashLight | uint64(l.Level)<<8 

 } 

  

 func (l LitPumpkin) Hash() uint64 { 

         return hashLitPumpkin | uint64(l.Facing)<<8 

 } 

  

 func (l Log) Hash() uint64 { 

         return hashLog | uint64(l.Wood.Uint8())<<8 | uint64(boolByte(l.Stripped))<<11 | uint64(l.Axis)<<12 

 } 

  

 func (Melon) Hash() uint64 { 

         return hashMelon 

 } 

  

 func (m MelonSeeds) Hash() uint64 { 

         return hashMelonSeeds | uint64(m.Growth)<<8 | uint64(m.Direction)<<16 

 } 

  

 func (MossCarpet) Hash() uint64 { 

         return hashMossCarpet 

 } 

  

 func (NetherBrickFence) Hash() uint64 { 

         return hashNetherBrickFence 

 } 

  

 func (n NetherBricks) Hash() uint64 { 

         return hashNetherBricks | uint64(n.Type.Uint8())<<8 

 } 

  

 func (NetherGoldOre) Hash() uint64 { 

         return hashNetherGoldOre 

 } 

  

 func (NetherQuartzOre) Hash() uint64 { 

         return hashNetherQuartzOre 

 } 

  

 func (NetherSprouts) Hash() uint64 { 

         return hashNetherSprouts 

 } 

  

 func (n NetherWart) Hash() uint64 { 

         return hashNetherWart | uint64(n.Age)<<8 

 } 

  

 func (NetheriteBlock) Hash() uint64 { 

         return hashNetheriteBlock 

 } 

  

 func (Netherrack) Hash() uint64 { 

         return hashNetherrack 

 } 

  

 func (NoteBlock) Hash() uint64 { 

         return hashNoteBlock 

 } 

  

 func (o Obsidian) Hash() uint64 { 

         return hashObsidian | uint64(boolByte(o.Crying))<<8 

 } 

  

 func (PackedIce) Hash() uint64 { 

         return hashPackedIce 

 } 

  

 func (p Planks) Hash() uint64 { 

         return hashPlanks | uint64(p.Wood.Uint8())<<8 

 } 

  

 func (Podzol) Hash() uint64 { 

         return hashPodzol 

 } 

  

 func (p Potato) Hash() uint64 { 

         return hashPotato | uint64(p.Growth)<<8 

 } 

  

 func (p Prismarine) Hash() uint64 { 

         return hashPrismarine | uint64(p.Type.Uint8())<<8 

 } 

  

 func (p Pumpkin) Hash() uint64 { 

         return hashPumpkin | uint64(boolByte(p.Carved))<<8 | uint64(p.Facing)<<9 

 } 

  

 func (p PumpkinSeeds) Hash() uint64 { 

         return hashPumpkinSeeds | uint64(p.Growth)<<8 | uint64(p.Direction)<<16 

 } 

  

 func (q Quartz) Hash() uint64 { 

         return hashQuartz | uint64(boolByte(q.Smooth))<<8 

 } 

  

 func (QuartzBricks) Hash() uint64 { 

         return hashQuartzBricks 

 } 

  

 func (q QuartzPillar) Hash() uint64 { 

         return hashQuartzPillar | uint64(q.Axis)<<8 

 } 

  

 func (RawCopperBlock) Hash() uint64 { 

         return hashRawCopperBlock 

 } 

  

 func (RawGoldBlock) Hash() uint64 { 

         return hashRawGoldBlock 

 } 

  

 func (RawIronBlock) Hash() uint64 { 

         return hashRawIronBlock 

 } 

  

 func (s Sand) Hash() uint64 { 

         return hashSand | uint64(boolByte(s.Red))<<8 

 } 

  

 func (s Sandstone) Hash() uint64 { 

         return hashSandstone | uint64(s.Type.Uint8())<<8 | uint64(boolByte(s.Red))<<10 

 } 

  

 func (s SandstoneStairs) Hash() uint64 { 

         return hashSandstoneStairs | uint64(boolByte(s.Smooth))<<8 | uint64(boolByte(s.Red))<<9 | uint64(boolByte(s.UpsideDown))<<10 | uint64(s.Facing)<<11 

 } 

  

 func (SeaLantern) Hash() uint64 { 

         return hashSeaLantern 

 } 

  

 func (s SeaPickle) Hash() uint64 { 

         return hashSeaPickle | uint64(s.AdditionalCount)<<8 | uint64(boolByte(s.Dead))<<16 

 } 

  

 func (Shroomlight) Hash() uint64 { 

         return hashShroomlight 

 } 

  

 func (s Sign) Hash() uint64 { 

         return hashSign | uint64(s.Wood.Uint8())<<8 | uint64(s.Attach.Uint8())<<11 

 } 

  

 func (Snow) Hash() uint64 { 

         return hashSnow 

 } 

  

 func (SoulSand) Hash() uint64 { 

         return hashSoulSand 

 } 

  

 func (SoulSoil) Hash() uint64 { 

         return hashSoulSoil 

 } 

  

 func (s Sponge) Hash() uint64 { 

         return hashSponge | uint64(boolByte(s.Wet))<<8 

 } 

  

 func (SporeBlossom) Hash() uint64 { 

         return hashSporeBlossom 

 } 

  

 func (g StainedGlass) Hash() uint64 { 

         return hashStainedGlass | uint64(g.Colour.Uint8())<<8 

 } 

  

 func (p StainedGlassPane) Hash() uint64 { 

         return hashStainedGlassPane | uint64(p.Colour.Uint8())<<8 

 } 

  

 func (t StainedTerracotta) Hash() uint64 { 

         return hashStainedTerracotta | uint64(t.Colour.Uint8())<<8 

 } 

  

 func (s Stone) Hash() uint64 { 

         return hashStone | uint64(boolByte(s.Smooth))<<8 

 } 

  

 func (c StoneBricks) Hash() uint64 { 

         return hashStoneBricks | uint64(c.Type.Uint8())<<8 

 } 

  

 func (g TallGrass) Hash() uint64 { 

         return hashTallGrass | uint64(g.Type.Uint8())<<8 

 } 

  

 func (Terracotta) Hash() uint64 { 

         return hashTerracotta 

 } 

  

 func (t Torch) Hash() uint64 { 

         return hashTorch | uint64(t.Facing)<<8 | uint64(t.Type.Uint8())<<11 

 } 

  

 func (Tuff) Hash() uint64 { 

         return hashTuff 

 } 

  

 func (w Water) Hash() uint64 { 

         return hashWater | uint64(boolByte(w.Still))<<8 | uint64(w.Depth)<<9 | uint64(boolByte(w.Falling))<<17 

 } 

  

 func (s WheatSeeds) Hash() uint64 { 

         return hashWheatSeeds | uint64(s.Growth)<<8 

 } 

  

 func (w Wood) Hash() uint64 { 

         return hashWood | uint64(w.Wood.Uint8())<<8 | uint64(boolByte(w.Stripped))<<11 | uint64(w.Axis)<<12 

 } 

  

 func (d WoodDoor) Hash() uint64 { 

         return hashWoodDoor | uint64(d.Wood.Uint8())<<8 | uint64(d.Facing)<<11 | uint64(boolByte(d.Open))<<13 | uint64(boolByte(d.Top))<<14 | uint64(boolByte(d.Right))<<15 

 } 

  

 func (w WoodFence) Hash() uint64 { 

         return hashWoodFence | uint64(w.Wood.Uint8())<<8 

 } 

  

 func (f WoodFenceGate) Hash() uint64 { 

         return hashWoodFenceGate | uint64(f.Wood.Uint8())<<8 | uint64(f.Facing)<<11 | uint64(boolByte(f.Open))<<13 | uint64(boolByte(f.Lowered))<<14 

 } 

  

 func (s WoodSlab) Hash() uint64 { 

         return hashWoodSlab | uint64(s.Wood.Uint8())<<8 | uint64(boolByte(s.Top))<<11 | uint64(boolByte(s.Double))<<12 

 } 

  

 func (s WoodStairs) Hash() uint64 { 

         return hashWoodStairs | uint64(s.Wood.Uint8())<<8 | uint64(boolByte(s.UpsideDown))<<11 | uint64(s.Facing)<<12 

 } 

  

 func (t WoodTrapdoor) Hash() uint64 { 

         return hashWoodTrapdoor | uint64(t.Wood.Uint8())<<8 | uint64(t.Facing)<<11 | uint64(boolByte(t.Open))<<13 | uint64(boolByte(t.Top))<<14 

 } 

  

 func (w Wool) Hash() uint64 { 

         return hashWool | uint64(w.Colour.Uint8())<<8 

 }