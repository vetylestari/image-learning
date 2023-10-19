from tensorflow.keras.preprocessing.image import ImageDataGenerator
import os

# Set the paths to your data directories
data_dir = 'data/raw_data'  # Directory where your raw data is stored
output_dir = 'data/'  # Directory where preprocessed data will be saved

# Define parameters for data preprocessing
image_size = (224, 224)  # Desired image size
batch_size = 32  # Batch size for data generators

# Create data generators for training, validation, and test sets
datagen = ImageDataGenerator(
    rescale=1.0/255,  # Rescale pixel values to [0, 1]
    rotation_range=15,  # Randomly rotate images within a 15-degree range
    width_shift_range=0.2,  # Randomly shift the width of images by up to 20%
    height_shift_range=0.2,  # Randomly shift the height of images by up to 20%
    shear_range=0.2,  # Apply shear transformations
    zoom_range=0.2,  # Apply random zoom to images
    horizontal_flip=True,  # Randomly flip images horizontally
    fill_mode='nearest'  # Fill newly created pixels after rotation or shifting
)

# Split data into training, validation, and test directories
train_dir = os.path.join(output_dir, 'training')
validation_dir = os.path.join(output_dir, 'validation')
test_dir = os.path.join(output_dir, 'testing')

# Create the necessary directories
os.makedirs(train_dir, exist_ok=True)
os.makedirs(validation_dir, exist_ok=True)
os.makedirs(test_dir, exist_ok=True)

# Generate data splits and preprocess images
splits = ['training', 'validation', 'testing']
split_ratios = [0.6, 0.2, 0.2]  # Adjust the ratios as needed

for split, ratio in zip(splits, split_ratios):
    split_dir = os.path.join(data_dir, split)
    num_samples = len(os.listdir(split_dir))
    num_images = int(ratio * num_samples)

    # Create data generators for the split
    generator = datagen.flow_from_directory(
        split_dir,
        target_size=image_size,
        batch_size=batch_size,
        class_mode='categorical',
        shuffle=True
    )

    # Save the preprocessed data to the corresponding directory
    output_split_dir = os.path.join(output_dir, split)
    os.makedirs(output_split_dir, exist_ok=True)
    for i in range(num_images // batch_size):
        batch = generator.next()
        images, labels = batch
        for j in range(len(images)):
            image = images[j]
            label = labels[j]
            image_filename = f'{i * batch_size + j}.jpg'
            image_path = os.path.join(output_split_dir, label, image_filename)
            os.makedirs(os.path.dirname(image_path), exist_ok=True)
            image *= 255  # Convert back to the original scale
            image = image.astype(int)
            image = image[..., ::-1]  # Convert from BGR to RGB
            image = image.astype('uint8')
            tf.keras.preprocessing.image.save_img(image_path, image)

print("Data preprocessing complete.")