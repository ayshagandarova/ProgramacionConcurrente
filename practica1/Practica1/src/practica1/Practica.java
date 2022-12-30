/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package practica1;
import java.util.concurrent.Semaphore;
/**
 *
 * @author ashook
 */
public class Practica {

    static final int MAX_ESTUDIANTS = 4;
    static final int ESTUDIANTS = 10;
    static final int RONDAS = 3;
    static Semaphore sEntrada = new Semaphore(1);
    static Semaphore sDirector = new Semaphore(0);
    static Semaphore sMutex = new Semaphore(1);
    static volatile int contEstudiants = 0;
    public static Estat estat;
    public enum Estat {
        FORA, 
        ESPERANT, 
        DINS
    }

    static final String[] noms5 = {
        "Pelayo",
        "Beltrán",
        "Cayetano",
        "Borja",
        "Jacobo"
    };
    
    static final String[] noms10 = {
        "Bosco",
        "Pelayo",
        "Cayetano",
        "Jacobo",
        "Tristán",
        "Beltrán",
        "Guzmán",
        "Froilán",
        "Borja",
        "Sancho"
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
            estudiants[i] = new Thread(new Estudiant(noms10[i]));
            estudiants[i].start();
        }

        for (int i=0; i< ESTUDIANTS; i++) {
            estudiants[i].join();
        }
        dir.join();
        System.out.println("SIMULACIÓ ACABADA");
    }
    
}
