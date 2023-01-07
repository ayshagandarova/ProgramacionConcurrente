-- AUTOR@S: ANTONIO PUJOL y AISHA GANDAROVA

with Ada.Text_IO; use Ada.Text_IO;
with def_monitor; use def_monitor;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;

procedure Restaurant is

   -- Variables globales
   PERSONES : constant Integer := 7;

   type persona is record
      nom: Unbounded_string;
      tipo: Integer;
   end record;
   type aStrings is array (0 .. (PERSONES*2-1)) of unbounded_string;
    
   MAX_CAPACITAT : constant Integer := 3;
   noms : constant aStrings := (
      to_unbounded_string("Tristán"), 
      to_unbounded_string("Pelayo"), 
      to_unbounded_string("Sancho"), 
      to_unbounded_string("Borja"), 
      to_unbounded_string("Bosco"), 
      to_unbounded_string("Guzmán"),
      to_unbounded_string("Froilán"),
      to_unbounded_string("Nicolás"), 
      to_unbounded_string("Jacobo"),
      to_unbounded_string("Rodrigo"), 
      to_unbounded_string("Gonzalo"), 
      to_unbounded_string("JoseMari"), 
      to_unbounded_string("Cayetano"), 
      to_unbounded_string("Leopoldo"));

 -- Tipo protegit para la SC
   maitre : monitor(NUM_SALONS, 3);

   -- Especificación de la tarea fumadores
   task type fumador is
      entry Start (Nom : in Unbounded_String);
   end fumador;

   -- Especificaci�n de la tarea no fumadores
   task type no_fumador is
      entry Start (Nom : in Unbounded_String);
   end no_fumador;

   task body fumador is 
      My_Nom : Unbounded_String;
      Salon : Integer;
   begin
      accept Start (Nom : in Unbounded_String) do
         My_Nom := Nom;
      end Start;

      Put_Line("BON DIA som en " & to_string(My_Nom) & " i sóm fumador");

          -- SECCI�N CR�TICA
      maitre.entrarFum(My_Nom, Salon);
      Put_Line("En " & to_string(My_Nom) & " diu: Prendré el menú del dia. Som al saló 1");
      delay 0.1;  -- lo que tarda en cruzar
      Put_Line("En " & to_string(My_Nom) & " diu: Ja he dinat, el compte per favor");
      Put_Line("En " & to_string(My_Nom) & " SE'N VA");
      maitre.sortirSalon(My_Nom, Salon);
   end fumador;

   task body no_fumador is 
      My_Nom : Unbounded_String;
      Salon : Integer;
   begin
      accept Start (Nom : in Unbounded_String) do
         My_Nom := Nom;
      end Start;

      Put_Line("BON DIA som en " & to_string(My_Nom) & " i sóm no fumador");

         -- SECCI�N CR�TICA
      maitre.entrarNoFum(My_Nom, Salon);
      Put_Line("En " & to_string(My_Nom) & " diu: Prendré el menú del dia. Som al saló 1");
      delay 0.1;  -- lo que tarda en cruzar
      Put_Line("En " & to_string(My_Nom) & " diu: Ja he dinat, el compte per favor");
      Put_Line("En " & to_string(My_Nom) & " SE'N VA");
      maitre.sortirSalon(My_Nom, Salon);
   end no_fumador;

     -- ARRAY DE TAREAS --
   type fumadors is array (0 .. PERSONES-1) of fumador;
   fum : fumadors;

   type no_fumadors is array (0 .. PERSONES-1) of no_fumador;
   no_fum : no_fumadors;

begin
   -- PROGRAMA PRINCIPAL --
   maitre.inicializarSalons;
   for i in 0 .. (PERSONES-1) loop
      fum(i).Start(noms(i));
      no_fum(i).Start(noms(i+PERSONES));
   end loop;

end Restaurant;