from tensorflow.keras.preprocessing.image import ImageDataGenerator

# Define path to the testing data
test_dir = 'data/testing'

# Define parameters for data preprocessing
image_size = (224, 224)
batch_size = 32

# Data preprocessing using ImageDataGenerator
test_datagen = ImageDataGenerator(rescale=1.0/255)

# Create a data generator for testing
test_generator = test_datagen.flow_from_directory(
    test_dir,
    target_size=image_size,
    batch_size=batch_size,
    class_mode='categorical'
)