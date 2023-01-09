-- AUTOR@S: ANTONIO PUJOL y AISHA GANDAROVA
-- video: https://www.dropbox.com/s/34q65upbsmyyr2c/Pr%C3%A1ctica%202%20Programaci%C3%B3n%20concurrente.mp4?dl=0

with Ada.Text_IO; use Ada.Text_IO;
with def_monitor; use def_monitor;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;

procedure Restaurant is

   -- Constantes
   PERSONES : constant Integer := 7;

   type aStrings is array (0 .. (PERSONES*2-1)) of unbounded_string;  

   NOMS : constant aStrings := (
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
      to_unbounded_string("Leopoldo")
   );

   -- Objeto protegido 
   maitre : monitor; 

   -- Tarea fumadores
   task type fumador is
      entry Start (Nom : in Unbounded_String);
   end fumador;

   -- Tarea no_fumadores
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

      -- sección crítica
      maitre.entrarFum(My_Nom, Salon);
      Put_Line("En " & to_string(My_Nom) & " diu: Prendré el menú del dia. Som al saló 1");
      delay 0.1;  -- lo que tarda en comer y fumar
      Put_Line("En " & to_string(My_Nom) & " diu: Ja he dinat, el compte per favor");
      Put_Line("En " & to_string(My_Nom) & " SE'N VA");
      maitre.sortirSalon(My_Nom, Salon);

   end fumador;

   -- Tarea fumador
   task body no_fumador is 
      My_Nom : Unbounded_String;
      Salon : Integer;
   begin
      accept Start (Nom : in Unbounded_String) do
         My_Nom := Nom;
      end Start;

      Put_Line("BON DIA som en " & to_string(My_Nom) & " i sóm no fumador");

      -- sección crítica
      maitre.entrarNoFum(My_Nom, Salon);
      Put_Line("En " & to_string(My_Nom) & " diu: Prendré el menú del dia. Som al saló 1");
      delay 0.1;  -- lo que tarda en comer
      Put_Line("En " & to_string(My_Nom) & " diu: Ja he dinat, el compte per favor");
      Put_Line("En " & to_string(My_Nom) & " SE'N VA");
      maitre.sortirSalon(My_Nom, Salon);
   end no_fumador;

   -- Array de fumadors
   type fumadors is array (0 .. PERSONES-1) of fumador;
   fum : fumadors;

   -- Array de no_fumadors
   type no_fumadors is array (0 .. PERSONES-1) of no_fumador;
   no_fum : no_fumadors;

begin
   maitre.inicializarSalons; -- inicializamos las variables del objeto protegido
   for i in 0 .. (PERSONES-1) loop -- lanzamos las tareas de los fumadores y no fumadores
      fum(i).Start(NOMS(i)); 
      no_fum(i).Start(NOMS(i+PERSONES));
   end loop;

end Restaurant;