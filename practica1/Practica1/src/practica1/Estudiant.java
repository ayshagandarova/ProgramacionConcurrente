package practica1;

public class Estudiant implements Runnable{
    private String name;
    public Estudiant(String name) {
        this.name = name;
    }

    @Override
    public void run() {
        try {
            Thread.sleep((long)(Math.random()*2000));
            
            Practica.sEntrada.acquire();
            Practica.sMutex.acquire();
            Practica.contEstudiants++;
            System.out.printf("%s entra a la sala d'estudi, nombre estudiants: %d\n",name, Practica.contEstudiants);
            

            if(Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){
                System.out.println(name + ": FESTA!!!!!");
                if(Practica.estat == Practica.Estat.ESPERANT){
                    Practica.sDirector.release();
                } 
                Practica.sMutex.release();
            }else{
                System.out.println(name + " estudia");
                Practica.sMutex.release();
            }
            Practica.sEntrada.release();
            
            Thread.sleep((long)(Math.random()*5000));

            Practica.sMutex.acquire();
            Practica.contEstudiants--;
            System.out.println(name + " surt de la sala d'estudi, nombre estudiants: " + Practica.contEstudiants);
            if (Practica.contEstudiants == 0 && Practica.estat != Practica.Estat.FORA){
                System.out.println(name + ": ADEU Senyor Director, pot entrar si vol, no hi ha ningú");
                Practica.sDirector.release(); // cuando llega a max y cuando llega a 0
            }
            Practica.sMutex.release();

        } catch (InterruptedException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }

}
