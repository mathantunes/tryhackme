import pickle
with open("download.dat", "rb") as file:
	pickle_data = file.read()
creds = pickle.loads(pickle_data)
print(creds)
