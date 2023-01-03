-- AUTOR@S: ANTONIO PUJOL y AISHA GANDAROVA

with Ada.Text_IO; use Ada.Text_IO;
with def_monitor; use def_monitor;

procedure Restaurant is

   -- Variables globales
   PERSONES : constant Integer := 7;

   -- Tipo protegit para la SC
   maitre : cordaMonitor;

   -- Especificación de la tarea fumadores
   task type fumador is
      entry Start (Idx : in Integer);
   end fumador;

   -- Especificaci�n de la tarea no fumadores
   task type no_fumador is
      entry Start (Idx : in Integer);
   end no_fumador;

   task body fumador is 
      My_Idx : Integer;
   begin
      accept Start (Idx : in Integer) do
         My_Idx := Idx;
      end Start;

      Put_Line
        ("BON DIA, soc en Babu� Nord" & Integer'Image (My_Idx) &
         " i vaig cap al Sud");

          -- SECCI�N CR�TICA
         monitor.nordEntrar;
         Put_Line
           ("Nord" & Integer'Image (My_Idx) &
            ": �s a la corda i travessa cap al Sud");
         delay 0.1;  -- lo que tarda en cruzar
         monitor.nordSortir;
   end fumador;

   task body no_fumador is 
      My_Idx : Integer;
   begin
      accept Start (Idx : in Integer) do
         My_Idx := Idx;
      end Start;

      Put_Line
        ("BON DIA, soc en Babu� Nord" & Integer'Image (My_Idx) &
         " i vaig cap al Sud");

          -- SECCI�N CR�TICA
         monitor.nordEntrar;
         Put_Line
           ("Nord" & Integer'Image (My_Idx) &
            ": �s a la corda i travessa cap al Sud");
         delay 0.1;  -- lo que tarda en cruzar
         monitor.nordSortir;
   end no_fumador;


   -- ARRAY DE TAREAS --
   type fumadors is array (1 .. PERSONES) of fumador;
   fum : fumadors;

   type no_fumadors is array (1 .. PERSONES) of no_fumador;
   no_fum : no_fumadors;

begin
   -- PROGRAMA PRINCIPAL --
   for Idx in 1 .. PERSONES loop
      fum (Idx).Start (Idx);
      no_fum (Idx).Start (Idx);
   end loop;

end Restaurants;