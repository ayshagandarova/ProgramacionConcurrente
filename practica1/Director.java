package practica1;


public class Director implements Runnable{
    static int rondas;
    static Estat estat;
    private enum Estat {
        FORA, 
        ESPERANT, 
        DINS
    }

    public Director(int rondas){
        this.rondas = rondas;
        estat = Estat.FORA;
    }

    public Estat getEstat(){
        return estat;
    }

    @Override
    public void run() {
        for (int i = 1; i<= rondas; i++){
            try {
                System.out.println("El Sr. Director comença la ronda");
                Practica.sMutex.acquire();
                if (Practica.contEstudiants == 0){
                    System.out.println("El Director veu que no hi ha ningú a la sala d'estudis");
                    Practica.sMutex.release();
                }else if (Practica.contEstudiants < Practica.MAX_ESTUDIANTS){
                    Practica.sMutex.release();
                    estat = Estat.ESPERANT;
                    System.out.println("El Director està esperant per entrar. No molesta als que estudien");
                    Practica.sDirector.acquire();
                    Practica.sMutex.acquire();
                    if (Practica.contEstudiants > Practica.MAX_ESTUDIANTS){
                        Practica.sMutex.release();
                        estat = Estat.DINS;
                        System.out.println("El Director està dins la sala d'estudi: S'HA ACABAT LA FESTA!");
                        Practica.sEntrada.acquire();// espera a que todos salgan
                        Practica.sDirector.acquire();
                    }else {
                        Practica.sMutex.release();
                        estat = Estat.FORA;
                        System.out.println("El Director veu que no hi ha ningú a la sala d'estudis");
                    }
                }else if (Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){
                    Practica.sMutex.release();
                    estat = Estat.DINS;
                    System.out.println("El Director està dins la sala d'estudi: S'HA ACABAT LA FESTA!");
                    Practica.sEntrada.acquire();// espera a que todos salgan
                    Practica.sDirector.acquire();
                }

                System.out.println("El Director acaba la ronda " + i + " de " + rondas);
                Thread.sleep((long)(Math.random()*5000));
            } catch (InterruptedException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
        }
    }
}
