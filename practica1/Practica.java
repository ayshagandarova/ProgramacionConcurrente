package practica1;

import java.util.concurrent.Semaphore;

public class Practica {
    static final int MAX_ESTUDIANTS = 4;
    static final int ESTUDIANTS = 5;
    static final int RONDAS = 3;
    static Semaphore sEntrada = new Semaphore(1);
    static Semaphore sDirector = new Semaphore(1);
    static Semaphore sMutex = new Semaphore(1);
    static volatile int contEstudiants = 0;
  


    static final String[] noms = {
        "Pelayo",
        "Beltrán",
        "Cayetano",
        "Borja",
        "Jacobo"
    };

    public static void main(String[] args) throws InterruptedException {
        Thread[] estudiants = new Thread[ESTUDIANTS];
        Thread dir;

        System.out.println("SIMULACIÓ DE LA SALA D'ESTUDI");
        System.out.println("Nombre total d'estudiants: " + ESTUDIANTS);
        System.out.println("Nombre màxim d'estudiants: " + MAX_ESTUDIANTS);
        

        dir = new Thread(new Director(RONDAS));
        dir.start();
        for (int i=0; i< ESTUDIANTS; i++) {
            estudiants[i] = new Thread(new Estudiant(noms[i]));
            estudiants[i].start();
        }

        for (int i=0; i< ESTUDIANTS; i++) {
            estudiants[i].join();
        }
        dir.join();
        System.out.println("SIMULACIÓ ACABADA");
    }
}