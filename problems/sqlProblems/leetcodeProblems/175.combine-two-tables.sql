# Write your MySQL query statement below
select person.firstName, person.lastName, address.city, address.state 
from person 
Left Join Address on person.personId=address.personId