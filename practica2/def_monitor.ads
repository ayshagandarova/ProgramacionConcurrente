with Ada.Text_IO; use Ada.Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;

package def_monitor is

   NUM_SALONS : constant Integer := 3;
   TIPUS_SALON : constant Integer := 3;
  type rSalons is record
    Capacitat: Integer;
    Tipus: Integer; -- fumadors=0 no_fumadors=1 2=ninguno
  end record;
  type aSalons is array (1 .. NUM_SALONS) of rSalons;
  type aStrings is array (1 ..TIPUS_SALON ) of unbounded_string;
  tipusPersones : constant aStrings := (to_unbounded_string("FUMADOR"), to_unbounded_string("NOFUMADOR"), to_unbounded_string("CAP"));
  
  protected type Monitor(numSalons:  Integer; maxCapacitat: Integer) is
    procedure inicializarSalons;
    entry entrarSalon (Nom : in Unbounded_String; Tipo : in Integer; Salon : out Integer);
    procedure sortirSalon (Nom : in Unbounded_String; Salon : in Integer);

    function asignarSalon(Nom: in Unbounded_String; Tipo : in Integer; IdSalon : out Integer) return Boolean;
  private
    NUM_SALONS : Integer := numSalons;
    salons: aSalons;
    tipusPers: aStrings:=tipusPersones;
    MAX_CAPACITAT : Integer := maxCapacitat;
  end Monitor;
end def_monitor;