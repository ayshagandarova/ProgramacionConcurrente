with Ada.Text_IO; use Ada.Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;

package def_monitor is

  NUM_SALONS : constant Integer := 3;
  TIPUS_SALON : constant Integer := 3;
  type rSalons is record
    Capacitat: Integer;
    Tipus: Integer; -- fumadors=0 no_fumadors=1 2=ninguno
  end record;
  type aSalons is array (0 .. NUM_SALONS-1) of rSalons;
  type aStrings is array (0 ..TIPUS_SALON-1) of unbounded_string;
  tipusPersones : constant aStrings := (to_unbounded_string("FUMADOR"), to_unbounded_string("NOFUMADOR"), to_unbounded_string("CAP"));
  
  type aBooleanEntrada is array (0..1) of Boolean;
  type aBooleanSalons is array (0..2) of Boolean;

  protected type Monitor(numSalons:  Integer; maxCapacitat: Integer) is
    procedure inicializarSalons;
    entry entrarFum (Nom : in Unbounded_String; Salon : out Integer);
    entry entrarNoFum (Nom : in Unbounded_String; Salon : out Integer);
    procedure sortirSalon (Nom : in Unbounded_String; Salon : in Integer);
    
   
    procedure asignarSalon(Nom: in Unbounded_String; Tipo : in Integer; IdSalon : out Integer);
  private
    potEntrar : aBooleanEntrada;
    dispoSalon : aBooleanSalons;
    NUM_SALONS : Integer := numSalons;
    salons: aSalons;
    tipusPers: aStrings:=tipusPersones;
    MAX_CAPACITAT : Integer := maxCapacitat;
  end Monitor;
end def_monitor;