package main

import "github.com/nsf/termbox-go"

const num_frames = 10

var frames = [num_frames]string{
`

                            .':loxxxxdoc,.
                         .:dc,,dkkkxolldkkl
                        :xkk.  .kk: ... .o'
                       okkkkl''ck: .oool  ocl:
                      ;kkkkkkkkkk. :oooo, ;kkk:
                     .xkkkkkkkkkk, ;oooo, ,kkkk.
                     dkkkkkkkkkkko .oooo. lkkkk;
                    :kkkkkkkkkkkkk: .ll. ;kkkkk.
                    :kkkkkkkkkkkkkko. . lkkkkkl
                 'd. okkkkkkkkkkkkkkx, okkkkkx
               .ckkx' ,dkkkkkkkkkkkkkkxkkkkkk.
             .:xkkkkko' .;ldxxkkkkkkkkkkkkkk:
            ;xkkkkkkkkkxl;....lkkkkkkkkkkkkk.
          ;xkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk'
        :xkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkl
       ,:::::::::::::::::::::::::::::::::::::.`,

`                           ...
                      .;oOKXXXKOo,
                      ;OXXXXXXXXXXd
                  .k.  ,XXKc....;O,
                 ;KXx;;xXK' ;ooc. xxxK;
                .KXXXXXXXx 'ooooc 'XXXX;
                dXXXXXXXXx 'ooool  KXXXk
               .KXXXXXXXXK. loooc .XXXXK
               :XXXXXXXXXXo .lol. xXXXXX.
               lXXXXXXXXXXXo  , .OXXXXXX.
               .KXXXXXXXXXXX0, .0XXXXXXX,
              . 'OXXXXXXXXXXXXxKXXXXXXXXo
              xO, ,dKXXXXXXXXXXXXXXXXXXXK;
             lXXXOl'..,;;kXXXXXXXXXXXXXXXXc
           ;OXXXXXXXKOxddOXXXXXXXXXXXXXXXXXo
         :0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXk.
       ,0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXd
       :llllllllllllllllllllllllllllllllllllll`,

`
                 .';;;;;,'.
              ;dk0KKKKK0OO0k:
            ;0K.  xKKo. .. ;O.
           oKKK.  cKl .loo: .,.
          :KKKK0kkKK' :oooo' lKO.
          OKKKKKKKKK, :oooo, :KKO;
         ,KKKKKKKKKKo .oooo. oKKKKO.
         oKKKKKKKKKKK; ,oo' ;KKKKKKx
         xKKKKKKKKKKKKl .. oKKKKKKKo
         :KKKKKKKKKKKKKk' dKKKKKKKKd
          dKKKKKKKKKKKKKKOKKKKKKKKKK,
           :0KKKKKKKKKKKKKKKKKKKKKKKKx,
          .. :x0KKKKKKKKKKKKKKKKKKKKKKK0x:.
          kKk:....'xKKKKKKKKKKKKKKKKKKKKKK0c
        .xKKKKKKOkk0KKKKKKKKKKKKKKKKKKKKKKKKk.
       '0KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKo
       :llllllllllllllllllllllllllllllllllllll`,

`

           ,ldxxxxxdl:,.
         c0KXXXXXXXXXXXXO:
       .Od  'KXXO,....:d..l.
      .0Xd   0X0 .coo:    dk
      kXXXOxOXXl ;oooo; :XXX0;
     lXXXXXXXXXl ;oooo: 'XXXXX:
    .KXXXXXXXXXO .oooo; :XXXXX0
    'XXXXXXXXXXXc 'oo: .0XXXXXO
    .XXXXXXXXXXXXl .' ,KXXXXXXo
     dXXXXXXXXXXXXO' :XXXXXXXXx
     .OXXXXXXXXXXXXXkXXXXXXXXXXo
       oXXXXXXXXXXXXXXXXXXXXXXXX0;
        .cOXXXXXXXXXXXXXXXXXXXXXXX0o'
       kx;..,clo0XXXXXXXXXXXXXXXXXXXXXOo;.
      .XXXX0xoccOXXXXXXXXXXXXXXXXXXXXXXXXXKo.
       lllllllllllllllllllllllllllllllllllllc`,

`

           .;cloollc;'.
        .lkkkkkkkkkkkkkx:
      .okdclkkkkdlclxkkxod.
     ,kkO.  'kd. ... ,k.  o.
    ckkkk:. ;k. cooo, 'o,;kx;
   ;kkkkkkkkko .ooooo  xkkkkko
   xkkkkkkkkkx  ooooo. xkkkkkk.
   kkkkkkkkkkk, ;ooo: .kkkkkkk.
   kkkkkkkkkkkk. ;o: .dkkkkkkx
   lkkkkkkkkkkkk; . 'kkkkkkkkx
   .xkkkkkkkkkkkko.,kkkkkkkkkkd.
    .lkkkkkkkkkkkkkkkkkkkkkkkkkkc.
    . .lkkkkkkkkkkkkkkkkkkkkkkkkkko;.
     o:. ':lookkkkkkkkkkkkkkkkkkkkkkkko:,.
     .xkxl;''.okkkkkkkkkkkkkkkkkkkkkkkkkkkkl.
      .::::::::::::::::::::::::::::::::::::::.`,

`


              .,;::::;,..
           .cxkkkkkkkkkkkd,
         .okl  .xkkkdc:;cxo
       .lkkkl   dkk; .,,. '.
      .xkkkkkdlokkl 'oooo. lkd.
     .kkkkkkkkkkkk, ;oooo: ,kkkx:
     ,kkkkkkkkkkkkc 'oooo: ,kkkkko
     .kkkkkkkkkkkkk. looo' lkkkkkx
      lkkkkkkkkkkkko  :l. :kkkkkkd
      ;kkkkkkkkkkkkkx,   lkkkkkkkl
       okkkkkkkkkkkkkkl.dkkkkkkkkx
      . ckkkkkkkkkkkkkkkkkkkkkkkkkd.
      k; .lkkkkkkkkkkkkkkkkkkkkkkkkkd;.
      kkx:. 'cdxkkkkkkkkkkkkkkkkkkkkkkkl,
      '::::,.    .::::::::::::::::::::::::,`,

`



                    .',;;;;;,'.
                 ,lkOOOOOOOOOOOkc
              .:kOOo;;xOOOOOOOOOOo
             cOOOOO.  'OOk:....ck.
           'kOOOOOOo'.lOk. :ll; .dcl:
          ,OOOOOOOOOOOOOc ,oooo; :OOOx.
          dOOOOOOOOOOOOO: ;ooooc 'OOOOd
          oOOOOOOOOOOOOOd .oooo: ;OOOOO.
          :OOOOOOOOOOOOOO; ;ooc  xOOOOO;
           oOOOOOOOOOOOOOO; .; .dOOOOOk.
         '' :OOOOOOOOOOOOOOo. 'kOOOOOd.
        .kOl..ckOOOOOOOOOOOOklOOOOOOo
       .kOOOOl. 'coxkkOOOOOOOOOOOOOOx.
       ,ccccccc:.     :ccccccccccccccc:;,..`,

`



                               ...
                         .;ldxxxxxxxdl;
                       ,oxxxxxxxxxxxxxxd.
                     ,dxxc  :xxxl.....:l
                   .dxxxxl  .xxc 'ooo;  ..
                  cxxxxxxxdoxxx. loooo. dxo
                 lxxxxxxxxxxxxx. loooo' lxxl
                'xxxxxxxxxxxxxx: ,oooo. dxxx'
                'xxxxxxxxxxxxxxx. :ol' :xxxx;
              .' cxxxxxxxxxxxxxxx, .. cxxxxo
            .lxx; ;xxxxxxxxxxxxxxxo..oxxxxx'
          'oxxxxxl. :xxxxxxxxxxxxxxxdxxxxxx'
        'oxxxxxxxxxl. .:ldxxxxxxxxxxxxxxxxxl
       .:::::::::::::;.     :cccccccccccccc:,`,

`



                               ..',,,'..
                            ':cldxxxxxollc,
                         .cdx.  'xxd,  .. '.
                       'lxxxx,  'xx; 'ooo:  .
                     'oxxxxxxxddxxx. loooo. lc
                    lxxxxxxxxxxxxxx. loooo' :x:
                   'xxxxxxxxxxxxxxx: 'oooo. oxx;
                 . .xxxxxxxxxxxxxxxd. ,ol. :xxxx.
                .d. cxxxxxxxxxxxxxxxd; .  cxxxxx'
               .oxd. ;dxxxxxxxxxxxxxxxo..oxxxxxd.
             .cxxxxx: .:dxxxxxxxxxxxxxxddxxxxxd.
           'lxxxxxxxxd:. .;cllxxxxxxxxxxxxxxxx,
         ,dxxxxxxxxxxxxxdc;'..oxxxxxxxxxxxxxxo
        ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,`,

`


                                .;clooooll:'.
                             'lc..:dddddlccod;
                          .:ddd'   odd:  .. .'
                         :dddddo;':dd: .oooc  c:
                       ,ddddddddddddd. :oooo, ,d.
                      cdddddddddddddd' ;oooo; ,dl
                     'dddddddddddddddl .oooo. :ddl
                     .dddddddddddddddd; .ll. 'dddd;
                  .o. cddddddddddddddddc  . :dddddc
                 .ldo. :ddddddddddddddddo. cdddddd;
                ,ddddd, .cddddddddddddddddoddddddc
             .;odddddddl, .':looddddddddddddddddc
           .cdddddddddddddl;....lddddddddddddddo
         ,odddddddddddddddddddddddddddddddddddd.
        ';;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;'`,
}

var colors = [num_frames]termbox.Attribute{
	// approx colors from original gif
	termbox.Attribute(210),   // peach
	termbox.Attribute(222),   // orange
	termbox.Attribute(120),   // green
	termbox.Attribute(123),   // cyan
	termbox.Attribute(111),   // blue
	termbox.Attribute(134),   // purple
	termbox.Attribute(177),   // pink
	termbox.Attribute(207),   // fuschia
	termbox.Attribute(206),   // magenta
	termbox.Attribute(204),   // red
}
