
<template>
  <div class="recipe-card">
    <div class="recipe-list">
      <div v-for="recipe in recipes" :key="recipe.id" class="recipe-item">
        <img :src="recipe.image" :alt="recipe.title" class="recipe-image">
        <div class="recipe-content">
          <h3>{{ recipe.title }}</h3>
          <p>{{ recipe.description }}</p>
          <div class="recipe-meta">
            <span><i class="fas fa-clock"></i> {{ recipe.cookTime }} mins</span>
            <span><i class="fas fa-user"></i> {{ recipe.servings }} servings</span>
          </div>
          <div class="recipe-tags">
            <span v-for="tag in recipe.tags" :key="tag" class="tag">{{ tag }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'RecipeCard',
  data() {
    return {
      recipes: [
        {
          id: 1,
          title: 'Spaghetti Carbonara',
          description: 'Classic Italian pasta dish with eggs, cheese, pancetta and black pepper',
          image: '/images/carbonara.jpg',
          cookTime: 30,
          servings: 4,
          tags: ['Italian', 'Pasta', 'Quick']
        },
        {
          id: 2,
          title: 'Chicken Curry',
          description: 'Aromatic Indian curry with tender chicken pieces',
          image: '/images/curry.jpg',
          cookTime: 45,
          servings: 6,
          tags: ['Indian', 'Spicy', 'Main Course']
        }
      ]
    }
  }
}
</script>

<style scoped>
.recipe-card {
  padding: 20px;
}

.recipe-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.recipe-item {
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.recipe-item:hover {
  transform: translateY(-5px);
}

.recipe-image {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.recipe-content {
  padding: 15px;
}

.recipe-content h3 {
  margin: 0 0 10px 0;
  color: #333;
}

.recipe-content p {
  color: #666;
  margin-bottom: 15px;
}

.recipe-meta {
  display: flex;
  gap: 15px;
  color: #777;
  margin-bottom: 10px;
}

.recipe-meta span {
  display: flex;
  align-items: center;
  gap: 5px;
}

.recipe-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.tag {
  background-color: #f0f0f0;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.9em;
  color: #555;
}
</style>
