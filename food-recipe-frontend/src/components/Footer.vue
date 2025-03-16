
<template>
  <footer class="footer">
    <div class="footer-content">
      <div class="footer-section">
        <h3>Recipe Hub</h3>
        <p>Your ultimate destination for delicious recipes and culinary inspiration.</p>
        <div class="social-links">
          <a href="#"><i class="fab fa-facebook"></i></a>
          <a href="#"><i class="fab fa-instagram"></i></a>
          <a href="#"><i class="fab fa-twitter"></i></a>
          <a href="#"><i class="fab fa-pinterest"></i></a>
        </div>
      </div>
      
      <div class="footer-section">
        <h4>Quick Links</h4>
        <ul>
          <li><router-link to="/">Home</router-link></li>
          <li><router-link to="/recipes">Recipes</router-link></li>
          <li><router-link to="/categories">Categories</router-link></li>
          <li><router-link to="/about">About Us</router-link></li>
        </ul>
      </div>

      <div class="footer-section">
        <h4>Newsletter</h4>
        <div class="newsletter-form">
          <input type="email" placeholder="Enter your email">
          <button>Subscribe</button>
        </div>
      </div>
    </div>
    
    <div class="footer-bottom">
      <p>© {{ currentYear }} Recipe Hub. All rights reserved.</p>
    </div>
  </footer>
</template>

<script>
export default {
  name: 'Footer',
  computed: {
    currentYear() {
      return new Date().getFullYear()
    }
  }
}
</script>

<style scoped>
.footer {
  background-color: #2c3e50;
  color: #ffffff;
  padding: 3rem 0 1rem 0;
  margin-top: 2rem;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  padding: 0 1rem;
}

.footer-section {
  margin-bottom: 1.5rem;
}

.footer-section h3 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
}

.footer-section h4 {
  font-size: 1.2rem;
  margin-bottom: 1rem;
}

.footer-section ul {
  list-style: none;
  padding: 0;
}

.footer-section ul li {
  margin-bottom: 0.5rem;
}

.footer-section a {
  color: #ffffff;
  text-decoration: none;
  transition: color 0.3s ease;
}

.footer-section a:hover {
  color: #3498db;
}

.social-links {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

.social-links a {
  font-size: 1.5rem;
  transition: transform 0.3s ease;
}

.social-links a:hover {
  transform: translateY(-3px);
}

.newsletter-form {
  display: flex;
  gap: 0.5rem;
}

.newsletter-form input {
  padding: 0.5rem;
  border: none;
  border-radius: 4px;
  flex: 1;
}

.newsletter-form button {
  padding: 0.5rem 1rem;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.newsletter-form button:hover {
  background-color: #2980b9;
}

.footer-bottom {
  text-align: center;
  margin-top: 2rem;
  padding-top: 1rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

@media (max-width: 768px) {
  .footer-content {
    grid-template-columns: 1fr;
  }
  
  .newsletter-form {
    flex-direction: column;
  }
  
  .newsletter-form button {
    width: 100%;
  }
}
</style>
