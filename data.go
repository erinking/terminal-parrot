package main

import "github.com/nsf/termbox-go"

const num_frames = 10

var frames = [num_frames]string{
`

                            .':loxxxxdoc,.
                         .:dc,,dkkkxolldkkl
                        :xkk.  .kk: ... .o'
                       okkkkl''ck: +eee+  ocl:
                      ;kkkkkkkkkk. +eeee+ ;kkk:
                     .xkkkkkkkkkk, +eeee+ ,kkkk.
                     dkkkkkkkkkkko +eeee+ lkkkk;
                    :kkkkkkkkkkkkk: +ee+ ;kkkkk.
                    :kkkkkkkkkkkkkko. + lkkkkkl
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
                 ;KXx;;xXK' +eee+ xxxK;
                .KXXXXXXXx +eeeee 'XXXX;
                dXXXXXXXXx +eeeee  KXXXk
               .KXXXXXXXXK. eeeee .XXXXK
               :XXXXXXXXXXo +eee+ xXXXXX.
               lXXXXXXXXXXXo  + .OXXXXXX.
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
           oKKK.  cKl +eee+ .,.
          :KKKK0kkKK' +eeee+ lKO.
          OKKKKKKKKK, +eeee+ :KKO;
         ,KKKKKKKKKKo +eeee+ oKKKKO.
         oKKKKKKKKKKK; +ee+ ;KKKKKKx
         xKKKKKKKKKKKKl ++ oKKKKKKKo
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
      .0Xd   0X0 +eee+    dk
      kXXXOxOXXl +eeee+ :XXX0;
     lXXXXXXXXXl +eeee+ 'XXXXX:
    .KXXXXXXXXXO +eeee+ :XXXXX0
    'XXXXXXXXXXXc +ee+ .0XXXXXO
    .XXXXXXXXXXXXl ++ ,KXXXXXXo
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
     ,kkO.  'kd. +++ ,k.  o.
    ckkkk:. ;k. eeee+ 'o,;kx;
   ;kkkkkkkkko +eeeee  xkkkkko
   xkkkkkkkkkx  eeeee+ xkkkkkk.
   kkkkkkkkkkk, +eee+ .kkkkkkk.
   kkkkkkkkkkkk. +e+ .dkkkkkkx
   lkkkkkkkkkkkk; + 'kkkkkkkkx
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
       .lkkkl   dkk; ++++ '.
      .xkkkkkdlokkl +eeee+ lkd.
     .kkkkkkkkkkkk, +eeee+ ,kkkx:
     ,kkkkkkkkkkkkc +eeee+ ,kkkkko
     .kkkkkkkkkkkkk. eeee+ lkkkkkx
      lkkkkkkkkkkkko  +e+ :kkkkkkd
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
           'kOOOOOOo'.lOk. +ee+ .dcl:
          ,OOOOOOOOOOOOOc +eeee+ :OOOx.
          dOOOOOOOOOOOOO: +eeeee 'OOOOd
          oOOOOOOOOOOOOOd +eeee+ ;OOOOO.
          :OOOOOOOOOOOOOO; +eee  xOOOOO;
           oOOOOOOOOOOOOOO; ++ .dOOOOOk.
         '' :OOOOOOOOOOOOOOo. 'kOOOOOd.
        .kOl..ckOOOOOOOOOOOOklOOOOOOo
       .kOOOOl. 'coxkkOOOOOOOOOOOOOOx.
       ,ccccccc:.     :ccccccccccccccc:;,..`,

`



                               ...
                         .;ldxxxxxxxdl;
                       ,oxxxxxxxxxxxxxxd.
                     ,dxxc  :xxxl.....:l
                   .dxxxxl  .xxc +eee+  ..
                  cxxxxxxxdoxxx. eeeee+ dxo
                 lxxxxxxxxxxxxx. eeeee+ lxxl
                'xxxxxxxxxxxxxx: +eeee+ dxxx'
                'xxxxxxxxxxxxxxx. +ee+ :xxxx;
              .' cxxxxxxxxxxxxxxx, ++ cxxxxo
            .lxx; ;xxxxxxxxxxxxxxxo..oxxxxx'
          'oxxxxxl. :xxxxxxxxxxxxxxxdxxxxxx'
        'oxxxxxxxxxl. .:ldxxxxxxxxxxxxxxxxxl
       .:::::::::::::;.     :cccccccccccccc:,`,

`



                               ..',,,'..
                            ':cldxxxxxollc,
                         .cdx.  'xxd,  ++ '.
                       'lxxxx,  'xx; +eee+  .
                     'oxxxxxxxddxxx. eeeee+ lc
                    lxxxxxxxxxxxxxx. eeeee+ :x:
                   'xxxxxxxxxxxxxxx: +eeee+ oxx;
                 . .xxxxxxxxxxxxxxxd. +ee+ :xxxx.
                .d. cxxxxxxxxxxxxxxxd; +  cxxxxx'
               .oxd. ;dxxxxxxxxxxxxxxxo..oxxxxxd.
             .cxxxxx: .:dxxxxxxxxxxxxxxddxxxxxd.
           'lxxxxxxxxd:. .;cllxxxxxxxxxxxxxxxx,
         ,dxxxxxxxxxxxxxdc;'..oxxxxxxxxxxxxxxo
        ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;,`,

`


                                .;clooooll:'.
                             'lc..:dddddlccod;
                          .:ddd'   odd:  ++ .'
                         :dddddo;':dd: +eeee  c:
                       ,ddddddddddddd. +eeee+ ,d.
                      cdddddddddddddd' +eeee+ ,dl
                     'dddddddddddddddl +eeee+ :ddl
                     .dddddddddddddddd; +ee+ 'dddd;
                  .o. cddddddddddddddddc  + :dddddc
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

var beak_color = termbox.Attribute(101)   // brown
