import tensorflow as tf

def preprocess_image(image):
    # Lakukan pra-pemrosesan gambar di sini, seperti normalisasi dan perubahan ukuran
    # Kembalikan gambar yang telah diproses
    return preprocessed_image

def preprocess_dataset(dataset):
    # Pra-pemrosesan seluruh dataset pelatihan dan pengujian
    preprocessed_dataset = dataset.map(preprocess_image)
    return preprocessed_dataset