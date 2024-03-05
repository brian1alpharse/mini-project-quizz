import React from "react";
import { ScrollView, StatusBar } from "react-native";

import animalQuestions from "../../data/animal";
import computerQuestions from "../../data/computers";

import { RowItem } from "../../components/RowItem";

export default ({ navigation }) => (
  <ScrollView>
    <StatusBar barStyle="dark-content" />
    <RowItem
      name="Hewan"
      color="#799496"
      onPress={() =>
        navigation.navigate("Quiz", {
          title: "Hewan",
          questions: animalQuestions,
          color: "#799496"
        })
      }
    />
    <RowItem
      name="Computers"
      color="#49475B"
      onPress={() =>
        navigation.navigate("Quiz", {
          title: "Computers",
          questions: computerQuestions,
          color: "#49475B"
        })
      }
    />
  </ScrollView>
);
