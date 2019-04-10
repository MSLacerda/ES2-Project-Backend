import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class TestJunit {
   @Test
	
   public void testAdd() {
      TestUserCode usr = new TestUserCode();
      assertEquals(10, usr.AddValor(5, 5));
   }

   public void testSub() {
      TestUserCode usr = new TestUserCode();
      assertEquals(0, usr.SubValor(5, 5));
   }
}