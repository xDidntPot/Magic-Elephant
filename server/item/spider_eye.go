 

 package item 

  

 import ( 

         "github.com/Pocketminer92/magic-alpaca/server/entity/effect" 

         "github.com/Pocketminer92/magic-alpaca/server/world" 

         "time" 

 ) 

  

 // SpiderEye is a poisonous food and brewing item. 

 type SpiderEye struct { 

         defaultFood 

 } 

  

 // Consume ... 

 func (SpiderEye) Consume(_ *world.World, c Consumer) Stack { 

         c.Saturate(2, 3.2) 

         c.AddEffect(effect.New(effect.Poison{}, 1, time.Second*5)) 

         return Stack{} 

 } 

  

 // EncodeItem ... 

 func (SpiderEye) EncodeItem() (name string, meta int16) { 

         return "minecraft:spider_eye", 0 

 }
