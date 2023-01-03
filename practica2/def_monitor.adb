with Ada.Text_IO; use Ada.Text_IO;
package body def_monitor is

   protected body monitor is

    -- Babuino del NORTE quiere entrar
    entry nordEntrar when numBabuinsSud = 0 and numBabuinsNord < 3 is
    begin
         numBabuinsNord := numBabuinsNord + 1;
         Put_Line("A la corda n'hi ha"
               & Integer'Image(numBabuinsNord) & " direcci� Sud");
    end nordEntrar;

    -- Babuino del NORTE quiere salir
    procedure nordSortir is
    begin
         numBabuinsNord := numBabuinsNord - 1;
         Put_Line("A la corda n'hi ha"
               & Integer'Image(numBabuinsNord) & " direcci� Sud");
    end nordSortir;

    -- Babuino del SUR quiere entrar
    entry sudEntrar when numBabuinsNord = 0 and numBabuinsSud < 3 is
    begin
         numBabuinsSud := numBabuinsSud + 1;
         Put_Line("     +++++A la corda n'hi ha"
               & Integer'Image(numBabuinsSud) & " direcci� Nord+++++");
    end sudEntrar;

    -- Babuino del SUR quiere salir
    procedure sudSortir is
    begin
         numBabuinsSud := numBabuinsSud - 1;
         Put_Line("     +++++A la corda n'hi ha" & Integer'Image(numBabuinsSud) & " direcci� Nord+++++");
    end sudSortir;

  end monitor;

end def_monitor;