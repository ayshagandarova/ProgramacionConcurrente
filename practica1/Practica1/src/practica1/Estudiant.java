package practica1;

public class Estudiant implements Runnable{
    private String name;
    public Estudiant(String name) {
        this.name = name;
    }

    @Override
    public void run() {
        try {
            Thread.sleep((long)(Math.random()*2000)); //El estudiant arriba a un moment random
            
            Practica.sEntrada.acquire(); //Bloquejam l'entrada y l'accés a variables
            Practica.sMutex.acquire();
            Practica.contEstudiants++;
            System.out.printf("%s entra a la sala d'estudi, nombre estudiants: %d\n",name, Practica.contEstudiants);
            
            //Segons la casuística actual de l'aula, farem una cosa o l'altra
            if(Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){ // Hi han suficients estudiants per fer festa
                System.out.println(name + ": FESTA!!!!!");
                if(Practica.estat == Practica.Estat.ESPERANT){ // Si el director ja hi era a l'aula, en seguida ho alliberam per acabar la festa
                    Practica.sDirector.release();
                } 
                Practica.sMutex.release();
            }else{ // No hi han suficients estudiants per fer festa, estudiam
                System.out.println(name + " estudia");
                Practica.sMutex.release();
            }
            Practica.sEntrada.release();
            
            Thread.sleep((long)(Math.random()*5000)); //temps d'estudi random

            Practica.sMutex.acquire();
            Practica.contEstudiants--; // El estudiant s'en va de l'aula
            System.out.println(name + " surt de la sala d'estudi, nombre estudiants: " + Practica.contEstudiants);
            
            //Si el estudiant és l'últim en sortir de l'aula
            if (Practica.contEstudiants == 0 && Practica.estat == Practica.Estat.ESPERANT){ // en cas de que no hi ha hagut festa
                System.out.println(name + ": ADEU Senyor Director, pot entrar si vol, no hi ha ningú");
                Practica.sDirector.release();
            } else if(Practica.contEstudiants == 0 && Practica.estat == Practica.Estat.DINS){ // cas en que hi ha hagut festa
                System.out.println(name + ": ADEU Senyor Director, es queda sol");
                Practica.sDirector.release();
            }
            Practica.sMutex.release();

        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

}
