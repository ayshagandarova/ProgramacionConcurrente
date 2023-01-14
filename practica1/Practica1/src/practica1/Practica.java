package practica1;
import java.util.concurrent.Semaphore;
/**
 *
 * @author Aisha Gandarova y Antonio Pujol
 * enllaç vídeo dropbox: https://www.dropbox.com/s/nc2x1jdk5h6lqi9/Pr%C3%A1ctica1.mp4?dl=0
 */
public class Practica {

    static final int MAX_ESTUDIANTS = 15; //Máxim d'estudiants permesos a l'aula d'estudi
    static final int ESTUDIANTS = 5; //Nombre d'estudiants al programa
    static final int RONDAS = 3; //Nombre de rondes que el director fará
    static Semaphore sEntrada = new Semaphore(1); //Semáfor per controlar l'entrada d'estudiants a l'aula d'estudi
    static Semaphore sDirector = new Semaphore(0); //Semáfor per controlar els bloquejos d'el director
    static Semaphore sMutex = new Semaphore(1); //Semáfor per controlar l'exclusió mutua de les variables consultades i modificades
    static volatile int contEstudiants = 0;
    public static Estat estat; //possible estat del director a un moment donat
    public enum Estat {
        FORA, 
        ESPERANT, 
        DINS
    }

    //valors d'estudiants per al cas de 5 alumnes al programa
    static final String[] noms5 = {
        "Pelayo",
        "Beltrán",
        "Cayetano",
        "Borja",
        "Jacobo"
    };
    
    //valors d'estudiants per al cas de 10 alumnes al programa
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
        
        dir = new Thread(new Director(RONDAS)); //Cream el fil director i el iniciam
        dir.start();
        for (int i=0; i< ESTUDIANTS; i++) { //Cream i iniciam els fils d'estudiants
            estudiants[i] = new Thread(new Estudiant(noms5[i]));
            estudiants[i].start();
        }

        for (int i=0; i< ESTUDIANTS; i++) {
            estudiants[i].join();
        }
        dir.join();
        System.out.println("SIMULACIÓ ACABADA");
    }
    
}
