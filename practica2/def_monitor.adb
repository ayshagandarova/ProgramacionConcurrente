with Ada.Text_IO; use Ada.Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;
package body def_monitor is
 protected body Monitor is
    entry entrarNoFum(Nom: in Unbounded_String; Salon : out Integer)  when  potEntrar(1) is 
          DispoSalon: Integer;
          IdSalon : Integer;
    begin
         
          asignarSalon(Nom, 1, IdSalon);

          DispoSalon:=MAX_CAPACITAT - salons(IdSalon).Capacitat;

          Salon := IdSalon + 1;
          Put_Line("******** En " & to_string(Nom) & " té taula al saló de no fumadors " & Salon'Image
          & ". Disponibilitat: " & DispoSalon'Image);
    end entrarNoFum;


    entry entrarFum(Nom: in Unbounded_String; Salon : out Integer)  when  potEntrar(0) is 
          DispoSalon: Integer;
          IdSalon : Integer;
    begin
          asignarSalon(Nom, 0, IdSalon);

          DispoSalon:=MAX_CAPACITAT - salons(IdSalon).Capacitat;

          Salon := IdSalon + 1;
          Put_Line("-------- En " & to_string(Nom) & " té taula al saló de fumadors " & Salon'Image
          & ". Disponibilitat: " & DispoSalon'Image);
    end entrarFum;

    

    procedure asignarSalon(Nom: in Unbounded_String; Tipo : in Integer; IdSalon : out Integer)  is
          encontrada : Boolean;
    begin
          encontrada :=false;
          for i in 0..(NUM_SALONS-1) loop
               if salons(i).Tipus = Tipo then 
                    if salons(i).Capacitat <  MAX_CAPACITAT then
                         if not encontrada then
                              encontrada := true;
                              salons(i).Capacitat := salons(i).Capacitat + 1;
                              IdSalon := i;
                         end if;
                    else
                         dispoSalon(i) := false;
                    end if;
               elsif salons(i).Tipus = 2 and not encontrada then
                    encontrada := true;
                    salons(i).Capacitat := salons(i).Capacitat + 1;
                    salons(i).Tipus :=  Tipo;
                    IdSalon := i;
               end if;
          end loop;

          
          encontrada :=false;
          potEntrar(0) := false;
          potEntrar(1) := false;
          for i in 0..(NUM_SALONS-1) loop
               if dispoSalon(i) then
                    if salons(i).Tipus = 2  then
                         potEntrar(0) := true;
                         potEntrar(1) := true;
                    else
                         potEntrar(salons(i).Tipus) := true;
                    end if;
               end if;
          end loop;

          
     end asignarSalon;

    procedure sortirSalon(Nom : in Unbounded_String; Salon : in Integer) is
          idSalon : Integer;
          disponibilidad : Integer;

    begin
          idSalon := Salon - 1 ;

          salons(idSalon).Capacitat := salons(idSalon).Capacitat - 1;

          potEntrar(salons(idSalon).Tipus) := true;
          dispoSalon(idSalon) := true;


          if salons(idSalon).Capacitat = 0 then
               salons(idSalon).Tipus := 2;
          end if;

          disponibilidad := MAX_CAPACITAT - salons(idSalon).Capacitat;
          if salons(idSalon).Tipus = 0 then --fumadors
               Put_Line("-------- En " & to_string(Nom) & " allibera una taula del saló " & Salon'Image & ". Disponibilitat: " & disponibilidad'Image & " Tipus: " & to_string(tipusPersones(salons(idSalon).Tipus)) );
          elsif salons(idSalon).Tipus = 1 then -- no fumadors
               Put_Line("******** En " & to_string(Nom) & " allibera una taula del saló " & Salon'Image & ". Disponibilitat: " & disponibilidad'Image & " Tipus: " & to_string(tipusPersones(salons(idSalon).Tipus)));
          end if;
    end sortirSalon;


   

     procedure inicializarSalons is
     begin
          potEntrar(0) := true;
          potEntrar(1) := true;

          dispoSalon(0) := true;
          dispoSalon(1) := true;
          dispoSalon(2) := true;

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