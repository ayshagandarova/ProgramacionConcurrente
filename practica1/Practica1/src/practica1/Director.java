package practica1;


public class Director implements Runnable{
    private static int rondas;


    public Director(int rondas){
        this.rondas = rondas;
        Practica.estat = Practica.estat.FORA; //Estat inicial del director
    }
    
    
    @Override
    public void run() {
        for (int i = 1; i<= rondas; i++){
            try {
                System.out.println("\t\tEl Sr. Director comença la ronda");
                Practica.sMutex.acquire();
                if (Practica.contEstudiants == 0){ // No hi ha ningú a l'aula
                    System.out.println("\t\tEl Director veu que no hi ha ningú a la sala d'estudis");
                    System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                    Practica.sMutex.release();
                    Thread.sleep((long)(Math.random()*1000)); // pausa rrandom fins a que fa la seguent ronda
                    continue;
                }else if (Practica.contEstudiants < Practica.MAX_ESTUDIANTS){ // hi han alumnes a l'aula però no hi ha festa
                    Practica.estat = Practica.Estat.ESPERANT;                    
                    System.out.println("\t\tEl Director està esperant per entrar. No molesta als que estudien");
                    Practica.sMutex.release();
                    Practica.sDirector.acquire(); //El director es queda bloquejat
                    
                    Practica.sMutex.acquire();
                    
                    //El director ha estat desbloquejat, mirarem si es per festa o que ja no hi ha ningú a l'aula
                    if (Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){ // Hi ha festa                        
                        Practica.estat = Practica.Estat.DINS; //Canviam l'estat
                        System.out.println("\t\tEl Director està dins la sala d'estudi: S'HA ACABAT LA FESTA!");
                        Practica.sEntrada.acquire(); //El director bloqueja que cap altre alumne entri a l'aula
                        Practica.sMutex.release();                        
                        Practica.sDirector.acquire(); //Queda bloquejat esperant a que surtin el alumnes de l'aula
                        System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                        
                        //resetejam l'estat del director
                        Practica.sMutex.acquire();
                        Practica.estat = Practica.Estat.FORA;
                        Practica.sMutex.release();
                        
                        Practica.sEntrada.release(); //Allibera lentrada a l'aula
                        Thread.sleep((long)(Math.random()*1000)); //Espera random fins a la seguent ronda
                        continue;
                    }else { //No hi ha festa, l'aula está buida                        
                        Practica.estat = Practica.Estat.FORA; //canviam estat
                        Practica.sMutex.release();
                        System.out.println("\t\tEl Director veu que no hi ha ningú a la sala d'estudis");
                    }
                }else if (Practica.contEstudiants >= Practica.MAX_ESTUDIANTS){ //El director arriba a veura que hi ha festa                    
                    Practica.estat = Practica.Estat.DINS;
                    System.out.println("\t\tEl Director està dins la sala d'estudi: S'HA ACABAT LA FESTA!");
                    Practica.sEntrada.acquire(); //El director bloqueja que cap altre alumne entri a l'aula
                    Practica.sMutex.release();
                    Practica.sDirector.acquire(); //Queda bloquejat esperant a que surtin el alumnes de l'aula
                    System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                    
                    //resetejam l'estat del director
                    Practica.sMutex.acquire();
                    Practica.estat = Practica.Estat.FORA;
                    Practica.sMutex.release();
                    
                    Practica.sEntrada.release();
                    Thread.sleep((long)(Math.random()*1000));
                    continue;
                }

                System.out.println("\t\tEl Director acaba la ronda " + i + " de " + rondas);
                
                Thread.sleep((long)(Math.random()*1000)); //Espera random fins a la seguent ronda
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
