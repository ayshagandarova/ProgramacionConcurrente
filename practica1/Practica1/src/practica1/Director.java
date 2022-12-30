package practica1;


public class Director implements Runnable{
    private static int rondas;


    public Director(int rondas){
        this.rondas = rondas;
        Practica.estat = Practica.estat.FORA;
    }
    
    
    @Override
    public void run() {
        for (int i = 1; i<= rondas; i++){
            try {
                System.out.println("\t\tEl Sr. Director comença la ronda");
                Practica.sMutex.acquire();
                if (Practica.contEstudiants == 0){
                    System.out.println("\t\tEl Director veu que no hi ha ningú a la sala d'estudis");
                    System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                    Practica.sMutex.release();
                    Thread.sleep((long)(Math.random()*1000));
                    continue;
                }else if (Practica.contEstudiants < Practica.MAX_ESTUDIANTS){
                    Practica.sMutex.release();
                    Practica.estat = Practica.Estat.ESPERANT;
                    System.out.println("\t\tEl Director està esperant per entrar. No molesta als que estudien");
                    Practica.sDirector.acquire();
                    Practica.sMutex.acquire();
                    if (Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){
                        Practica.sMutex.release();
                        Practica.estat = Practica.Estat.DINS;
                        System.out.println("\t\tEl Director està dins la sala d'estudi: S'HA ACABAT LA FESTA!");
                        Practica.sEntrada.acquire();// espera a que todos salgan
                        Practica.sDirector.acquire();
                        System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                        Practica.sEntrada.release();
                        Thread.sleep((long)(Math.random()*1000));
                        continue;
                    }else {
                        Practica.sMutex.release();
                        Practica.estat = Practica.Estat.FORA;
                        System.out.println("\t\tEl Director veu que no hi ha ningú a la sala d'estudis");
                    }
                }else if (Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){
                    Practica.sMutex.release();
                    Practica.estat = Practica.Estat.DINS;
                    System.out.println("\t\tEl Director està dins la sala d'estudi: S'HA ACABAT LA FESTA!");
                    Practica.sEntrada.acquire();// espera a que todos salgan
                    Practica.sDirector.acquire();
                    System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                    Practica.sEntrada.release();
                    Thread.sleep((long)(Math.random()*1000));
                    continue;
                }

                System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                Thread.sleep((long)(Math.random()*1000));
            } catch (InterruptedException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
        }
    }
}
