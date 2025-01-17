 

 package block 

  

 import ( 

         "github.com/Pocketminer92/magic-alpaca/server/block/cube" 

         "github.com/Pocketminer92/magic-alpaca/server/item" 

         "github.com/Pocketminer92/magic-alpaca/server/world" 

 ) 

  

 // Placer represents an entity that is able to place a block at a specific position in the world. 

 type Placer interface { 

         item.User 

         PlaceBlock(pos cube.Pos, b world.Block, ctx *item.UseContext) 

 }
