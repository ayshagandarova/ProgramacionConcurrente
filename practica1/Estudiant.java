package practica1;

public class Estudiant implements Runnable{
    private String name;
    public Estudiant(String name) {
        this.name = name;
    }

    @Override
    public void run() {
        try {
            Practica.sMutex.acquire();
            Practica.contEstudiants++;
            System.out.printf("%s entra a la sala d'estudi, nombre estudiants: %d\n",name, Practica.contEstudiants);
            

            if(Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){
                
                System.out.println(name + ": FESTA!!!!!");
                if(director.getEstat())
                Practica.sMutex.release();
                Practica.sDirector.release();
            }else{
                System.out.println(name + " estudia");
                Practica.sMutex.release();
            }
            Thread.sleep((long)(Math.random()*5000));

            Practica.sMutex.acquire();
            Practica.contEstudiants--;
            System.out.println(name + " surt de la sala d'estudi, nombre estudiants: " + Practica.contEstudiants);
            if (Practica.contEstudiants == 0){
                Practica.sDirector.release(); // cuando llega a max y cuando llega a 0
            }
            Practica.sMutex.release();

        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

}
