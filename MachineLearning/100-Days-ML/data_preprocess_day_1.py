import numpy as np
import pandas as pd

# importing data sets
print("-----------------------------Importing data sets-------------------------------------")
dataset = pd.read_csv('Data.csv')
X = dataset.iloc[:, :-1].values  # select all the column except the last one
Y = dataset.iloc[:, 3].values  # pick up the last column

print(X)
print(Y)

# handling missing values
print("--------------------------Handling missing values----------------------------------------")
from sklearn.impute import SimpleImputer
from sklearn.preprocessing import OneHotEncoder, LabelEncoder
from sklearn.compose import ColumnTransformer

imputer = SimpleImputer(missing_values=np.NaN, strategy="mean")
imputer = imputer.fit(X[:, 1:3])  # fit the Imputer to the three columns
X[:, 1:3] = imputer.transform(X[:, 1:3])
print(X)

# encoding categorical data
print("------------------------------Encoding categorical data------------------------------------")
label_encoder_X = LabelEncoder()
X[:, 0] = label_encoder_X.fit_transform(X[:, 0])
print(X)

print("------------------------------creating dummy variable------------------------------------")
ct = ColumnTransformer([("Country", OneHotEncoder(), [0])], remainder="passthrough")
X = ct.fit_transform(X)
label_encoder_Y = LabelEncoder()
Y = label_encoder_Y.fit_transform(Y)
print(X)
print(Y)

print("-------------------------------splitting datasets into training and test-----------------------------------")
from sklearn.model_selection import train_test_split

x_train, x_test, y_train, y_test = train_test_split(X, Y, test_size=0.2, random_state=0)
print("------------------x_train------------------")
print(x_train)
print("------------------x_test------------------")
print(x_test)
print("------------------y_train------------------")
print(y_train)
print("------------------y_test------------------")
print(y_test)

print("-----------------------------feature scaling------------------------------")
from sklearn.preprocessing import StandardScaler

sc_x = StandardScaler()
x_train = sc_x.fit_transform(x_train)
x_test = sc_x.fit_transform(x_test)
print("------- after scaling x_train--------")
print(x_train)
print("------- after scaling x_test--------")
print(x_test)
