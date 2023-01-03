with Ada.Text_IO; use Ada.Text_IO;
package def_monitor is

  protected type monitor is
    entry nordEntrar;
    procedure nordSortir;

    entry sudEntrar;
    procedure sudSortir;
  private

    numBabuinsNord : Integer := 0; -- Contador de Babuinos Norte
    numBabuinsSud  : Integer := 0; -- Contador de Babuinos Sur

  end monitor;

end def_monitor;