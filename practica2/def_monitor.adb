with Ada.Text_IO; use Ada.Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;
package body def_monitor is
 protected body Monitor is
    entry entrarSalon(Nom: in Unbounded_String; Tipo : in Integer; Salon : out Integer)  when  asignarSalon(Nom, Tipo, IdSalon)  is
          DispoSalon: Integer;
    begin
          salons(IdSalon).Capacitat := salons(IdSalon).Capacitat + 1;
          salons(IdSalon).Tipus :=  Tipo;
          DispoSalon:=MAX_CAPACITAT - salons(IdSalon).Capacitat;

          Salon := IdSalon + 1;
          if Tipo = 0 then --fumadors
               Put_Line("-------- En " & to_string(Nom) & "té taula al saló de fumadors " & IdSalon'Image
               & ". Disponibilitat: " & DispoSalon'Image);
          elsif Tipo = 1 then -- no fumadors
               Put_Line("******** En " & to_string(Nom) & "té taula al saló de no fumadors " & IdSalon'Image
               & ". Disponibilitat: " & DispoSalon'Image);
          end if;
    end entrarSalon;

    function asignarSalon(Nom: in Unbounded_String; Tipo : in Integer; IdSalon : out Integer) return Boolean is
    begin
          for i in 0..(NUM_SALONS-1) loop
               if salons(i).Tipus = Tipo then
                    if salons(i).Capacitat <  MAX_CAPACITAT then
                         IdSalon := i;
                         return true;
                    end if;
               elsif salons(i).Tipus = 2 then
                    IdSalon := i;
                    return true;
               end if;
          end loop;
          return false;
          
     end asignarSalon;
    procedure sortirSalon(Nom : in Unbounded_String; Salon : in Integer) is
          i : Integer;
          disponibilidad : Integer;

    begin
          i := Salon - 1 ;

          salons(i).Capacitat := salons(i).Capacitat - 1;

          if salons(i).Capacitat = 0 then
               salons(i).Tipus := 2;
          end if;

          disponibilidad := MAX_CAPACITAT - salons(i).Capacitat;
          if salons(i).Tipus = 0 then --fumadors
               Put_Line("-------- En " & to_string(Nom) & " allibera una taula del saló " & IdSalon'Image & ". Disponibilitat: " & disponibilidad'Image & " Tipus: " & to_string(tipusPersones(salons(i).Tipus)) );
          elsif salons(i).Tipus = 1 then -- no fumadors
               Put_Line("******** En " & to_string(Nom) & " allibera una taula del saló " & IdSalon'Image & ". Disponibilitat: " & disponibilidad'Image & " Tipus: " & to_string(tipusPersones(salons(i).Tipus)));
          end if;
    end sortirSalon;


   

     procedure inicializarSalons is
     begin

          for i in salons'Range loop
               salons(i).Capacitat := 0;
               salons(i).Tipus := 2;
          end loop;
          Put_Line("++++++++ El Maître està preparat");
          Put_Line("++++++++ Hi ha " & NUM_SALONS'Image & " salons amb capacitat de "
          & MAX_CAPACITAT'Image &" comensals cada un");
  
     end inicializarSalons;

  end Monitor;

end def_monitor;